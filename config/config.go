package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
)

var (
	Domain           string
	Protocol         string
	Host             string
	DBConfig         Database
	Secret           string
	Hibernations     string
	HTTPServerConfig HTTPServer
)

// loadEnv loads our environment variables
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

// Set configures our application
func Set() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&Domain, "domain", "localhost:8080", "Define the domain for our service to be hosted on")
	flag.StringVar(&Protocol, "protocol", "http://", "Define the protocol being used for our service (include ://)")
	flag.StringVar(&Hibernations, "hibernations", "../database/migrations/", "Define the path of SQL hibernations")
	flag.StringVar(&Secret, "secret", "EkFhJqLdPL7dCA4A", "Defines the header secret required to communicate with endpoints")
	Host = Protocol + Domain

	defer setHTTPServerConfig()
	defer setDBConfig()
}
