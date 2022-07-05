package config

import (
	"flag"
)

// Database is our repo configuration for establishing our connection
type Database struct {
	User    string `binding:"required"`
	Pass    string `binding:"required"`
	Host    string `binding:"required"`
	Port    string `binding:"required"`
	DbType  string `binding:"required"`
	DbName  string `binding:"required"`
	SSLMode string `binding:"required"`
}

var (
	Domain   string
	DBConfig Database
)

// Set configures our application
func Set() {
	flag.StringVar(&Domain, "domain", "localhost:8080", "Define the domain for our service to be hosted on")

	// TODO replace with envvar
	DBConfig = Database{
		User:    "postgres",
		Pass:    "postgres",
		Host:    "localhost",
		Port:    "5432",
		DbName:  "",
		DbType:  "postgres",
		SSLMode: "disable",
	}
}
