package server

import (
	"context"
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"httpgopha/cmd/server/ssl"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func HttpsServer(ip, port, dir, user, pass string, auth bool) {

	cert, key, err := ssl.KeyPair()
	if err != nil {
		log.Fatalln(err)
	}

	if !auth {

		e := echo.New()
		e.HideBanner = true
		fs := http.FileServer(http.Dir(dir))
		e.GET("/*", echo.WrapHandler(http.StripPrefix("/", fs)))

		// Start server
		go func() {
			e.Logger.Info("Starting HTTPS server on port %s", port)
			if err := e.StartTLS(ip+":"+port, cert, key); err != nil {
				e.Logger.Info("Shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}

	} else {

		e := echo.New()
		e.HideBanner = true
		fs := http.FileServer(http.Dir(dir))
		e.GET("/*", echo.WrapHandler(http.StripPrefix("/", fs)))
		e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) { // Auth
			// Be careful to use constant time comparison to prevent timing attacks
			if subtle.ConstantTimeCompare([]byte(username), []byte(user)) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(pass)) == 1 {
				return true, nil
			}
			return false, nil
		}))
		e.Use(middleware.Logger())

		// Start server
		go func() {
			e.Logger.Info("Starting HTTPS server with authentication on port %s", port)
			if err := e.StartTLS(ip+":"+port, cert, key); err != nil {
				e.Logger.Info("Shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}

	}
}
