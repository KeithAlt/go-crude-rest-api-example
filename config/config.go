package config

import (
	"flag"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
)

var (
	Domain         string
	DatabaseConfig repository.DBConfig
)

// Set configures our application
func Set() {
	flag.StringVar(&Domain, "domain", "localhost:8080", "Define the domain for our service to be hosted on")

	// TODO replace with envvar
	DatabaseConfig = repository.DBConfig{
		User:    "postgres",
		Pass:    "postgres",
		Host:    "localhost",
		Port:    "5432",
		DbName:  "",
		DbType:  "postgres",
		SSLMode: "disable",
	}
}
