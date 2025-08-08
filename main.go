package main

import (
	"flag"
	"log"

	"github.com/armon/go-socks5"
)

func main() {
	// Command-line flags
	addr := flag.String("addr", ":1080", "SOCKS5 proxy listen address")
	username := flag.String("user", "", "SOCKS5 proxy username (optional)")
	password := flag.String("pass", "", "SOCKS5 proxy password (optional)")
	flag.Parse()

	// Server configuration
	conf := &socks5.Config{}

	// If username and password are provided, set up authentication
	if *username != "" && *password != "" {
		creds := socks5.StaticCredentials{
			*username: *password,
		}
		cator := socks5.UserPassAuthenticator{Credentials: creds}
		conf.AuthMethods = []socks5.Authenticator{cator}
	}

	// Create a new SOCKS5 server
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	// Start listening
	log.Printf("Starting SOCKS5 proxy on %s", *addr)
	if err := server.ListenAndServe("tcp", *addr); err != nil {
		log.Fatal(err)
	}
}
