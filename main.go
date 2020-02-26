package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"httpgopha/cmd"
	"httpgopha/cmd/server"
	"os"
)

func main() {

	// Create new parser object
	parser := argparse.NewParser("httpgopha", "Quick webserver written in Go")

	// Create arguments
	var i = parser.String("i", "ip", &argparse.Options{Help: "IP [Default 0.0.0.0]"})
	var p = parser.String("p", "port", &argparse.Options{Help: "Port to server [Default 9090]"})
	var s = parser.Flag("s", "ssl", &argparse.Options{Help: "SSL [Default false]"})
	var d = parser.String("d", "directory", &argparse.Options{Help: "Directory"})
	var auth = parser.String("a", "auth", &argparse.Options{Help: "Authentication - Set username:password"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	// Send flags for logic check (valid ip/port, user/pass)
	a := cmd.ProcessArguments(i, p, d, auth, s)

	/* Debug flags
	fmt.Printf("IP: %v\n", a.Ip)
	fmt.Printf("SSL: %v\n", a.Ssl)
	fmt.Printf("Port: %v\n", a.Port)
	fmt.Printf("Directory: %v\n", a.Directory)
	fmt.Printf("Username: %v\n", a.Username)
	fmt.Printf("Password: %v\n", a.Password)
	fmt.Printf("Authentication: %v\n", a.Authentication)
	*/

	// Start webserver
	if a.Ssl {
		server.HttpsServer(a.Ip, a.Port, a.Directory, a.Username, a.Password, a.Authentication)
	} else {
		server.HttpServer(a.Ip, a.Port, a.Directory, a.Username, a.Password, a.Authentication)
	}
}
