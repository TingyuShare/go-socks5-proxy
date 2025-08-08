package mobile

import (
	"log"

	"github.com/armon/go-socks5"
)

// StartProxy starts the SOCKS5 proxy on the given address.
// This function is intended to be called from a native mobile application.
func StartProxy(addr, username, password string) {
	// Server configuration
	conf := &socks5.Config{}

	// If username and password are provided, set up authentication
	if username != "" && password != "" {
		creds := socks5.StaticCredentials{
			username: password,
		}
		cator := socks5.UserPassAuthenticator{Credentials: creds}
		conf.AuthMethods = []socks5.Authenticator{cator}
	}

	// Create a new SOCKS5 server
	server, err := socks5.New(conf)
	if err != nil {
		log.Printf("Failed to create SOCKS5 server: %v", err)
		return
	}

	// Start listening
	log.Printf("Starting SOCKS5 proxy on %s", addr)
	if err := server.ListenAndServe("tcp", addr); err != nil {
		log.Printf("SOCKS5 proxy failed: %v", err)
	}
}
