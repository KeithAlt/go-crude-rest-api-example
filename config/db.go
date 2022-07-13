package config

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
		User:    "postgres",
		Pass:    "postgres",
		Host:    "localhost",
		Port:    "5432",
		DbName:  "",
		DbType:  "postgres",
		SSLMode: "disable",
	}
}
