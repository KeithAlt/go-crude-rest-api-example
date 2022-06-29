package postgres

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qustavo/dotsql"
)

// DBConfig is our database configuration for establishing our connection
type DBConfig struct {
	User    string `binding:"required"`
	Pass    string `binding:"required"`
	Host    string `binding:"required"`
	Port    string `binding:"required"`
	DbType  string `binding:"required"`
	DbName  string `binding:"required"`
	SSLMode string `binding:"required"`
}

// Client is our connection instance
type Client struct {
	database *sql.DB
}

// NewClient creates a Client instance for us with values provided through a DBConfig
func NewClient(c DBConfig) (*Client, error) {
	psqlInfo := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		c.DbType, c.User, c.Pass, c.Host, c.Port, c.DbName, c.SSLMode,
	)

	db, err := sql.Open(c.DbType, psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &Client{db}, nil
}

// Insert inserts a new row into our database
func (c *Client) Insert(ctx *gin.Context, data interface{}) (interface{}, error) {
	// TODO: implement SQL
	return nil, nil
}

// UpdateById updates a pre-existing row by ID
func (c *Client) UpdateById(ctx *gin.Context, id, username, password, job string) error {
	// TODO: implement SQL
	return nil
}

// FindById finds a user by their id
func (c *Client) FindById(ctx *gin.Context, id string) (*sql.Row, error) {
	// TODO: implement SQL
	// TODO: Convert sql.Row to model object
	return nil, nil
}

// FindAll returns all rows of users
func (c *Client) FindAll(ctx *gin.Context) (*sql.Rows, error) {
	// TODO: implement SQL
	return nil, nil
}

// CreateTables will run 'create if not exist' statement migrations with use of dotsql
func (c *Client) CreateTables() error {
	dot, err := dotsql.LoadFromFile("./database/migrations/create_tables.sql")
	if err != nil {
		return fmt.Errorf("failed to load sql scripts: %w", err)
	}

	_, err = dot.Exec(c.database, "create-uuid-extension")
	if err != nil {
		return fmt.Errorf("failed to execute create extension migration: %w", err)
	}

	_, err = dot.Exec(c.database, "create-products-table")
	if err != nil {
		return fmt.Errorf("failed to execute create tables migration: %w", err)
	}
	return nil
}

// DropTables will run 'drop if not exist' statement migrations with use of dotsql
func (c *Client) DropTables() error {
	dot, err := dotsql.LoadFromFile("./database/migrations/create_tables.sql")
	if err != nil {
		return fmt.Errorf("failed to load sql scripts: %w", err)
	}

	_, err = dot.Exec(c.database, "drop-uuid-extension")
	if err != nil {
		return fmt.Errorf("failed to execute drop extension migration: %w", err)
	}

	_, err = dot.Exec(c.database, "drop-products-table")
	if err != nil {
		return fmt.Errorf("failed to execute drop tables migration: %w", err)
	}
	return nil
}
