package config

import "os"

// XXX All configs are pending gotdotenv implementation

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

// setDBConfig sets our database config
func setDBConfig() {
	DBConfig = Database{
		User:    os.Getenv("DATABASE_USERNAME"),
		Pass:    os.Getenv("DATABASE_PASSWORD"),
		Host:    os.Getenv("DATABASE_HOST"),
		Port:    os.Getenv("DATABASE_PORT"),
		DbName:  os.Getenv("DATABASE_NAME"),
		DbType:  os.Getenv("DATABASE_TYPE"),
		SSLMode: os.Getenv("DATABASE_SSLMODE"),
	}
}
