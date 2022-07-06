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
	Domain       string
	Protocol     string
	DBConfig     Database
	Hibernations string
)

// Set configures our application
func Set() {
	flag.StringVar(&Domain, "domain", "localhost:8080", "Define the domain for our service to be hosted on")
	flag.StringVar(&Protocol, "protocol", "http://", "Define the protocol being used for our service (include ://)")
	flag.StringVar(&Hibernations, "hibernations", "../database/migrations/", "Define the path of SQL hibernations")

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
