package config

import (
	"flag"
)

// XXX All configs are pending gotdotenv implementation

var (
	Domain           string
	Protocol         string
	Host             string
	DBConfig         Database
	Secret           string
	Hibernations     string
	HTTPServerConfig HTTPServer
)

// Set configures our application
func Set() {
	flag.StringVar(&Domain, "domain", "localhost:8080", "Define the domain for our service to be hosted on")
	flag.StringVar(&Protocol, "protocol", "http://", "Define the protocol being used for our service (include ://)")
	flag.StringVar(&Hibernations, "hibernations", "../database/migrations/", "Define the path of SQL hibernations")
	flag.StringVar(&Secret, "secret", "EkFhJqLdPL7dCA4A", "Defines the header secret required to communicate with endpoints")
	Host = Protocol + Domain

	setHTTPServerConfig()
	setDBConfig()
}
