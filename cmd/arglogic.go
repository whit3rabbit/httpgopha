package cmd

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Arguments struct {
	Ip             string
	Port           string
	Directory      string
	Username       string
	Password       string
	Ssl            bool
	Authentication bool
}

// Validate IP address
func checkAddress(host string) bool {
	return net.ParseIP(host) != nil
}

// Validate Port
func checkIfPort(str string) bool {
	if i, err := strconv.Atoi(str); err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

func ProcessArguments(ip, port, directory, auth *string, ssl *bool) *Arguments {

	a := new(Arguments) // Create new struct for Arguments

	DefaultIP := "0.0.0.0"
	DefaultPort := "9090"

	// If user didn't supply, then use default IP
	if len(*ip) > 0 {
		a.Ip = *ip
	} else {
		a.Ip = DefaultIP
	}

	// If user didn't supply, then use default port
	if len(*port) > 0 {
		a.Port = *port
	} else {
		a.Port = DefaultPort
	}

	// Validate IP
	isIPValid := checkAddress(a.Ip)
	if !isIPValid {
		fmt.Printf("[!] Ipv4 Address is invalid!")
		os.Exit(1)
	}

	// Validate port number
	isPortValid := checkIfPort(a.Port)
	if !isPortValid {
		fmt.Printf("[!] Port number is invalid!")
		os.Exit(1)
	}

	// Does user want SSL?
	if *ssl {
		a.Ssl = *ssl
	} else {
		a.Ssl = false
	}

	// If user didn't supply, then use default port
	if len(*directory) > 0 {
		a.Directory = *directory
	} else {
		a.Directory = "."
	}

	// I
	a.Authentication = false
	if len(*auth) > 0 {
		// Split username and password
		split := strings.Split(*auth, ":")
		username := split[0]
		password := split[1]

		// Set authentication
		a.Authentication = true
		a.Username = username
		a.Password = password

	}

	return a

}
