package postgres

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
	"github.com/rocketlaunchr/dbq/v2"
	"log"
	"time"
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

// productModel defines our model
type productModel struct {
	Name        string
	Price       int
	Description string
	CreatedAt   string
	UpdatedAt   string
}

// productQuery defines our model for communicating with our database
type productQueryModel struct {
	ID          string `dbq:"id"`
	GUID        string `dbq:"guid"`
	Name        string `dbq:"name"`
	Price       string `dbq:"price"`
	Description string `dbq:"description"`
	CreatedAt   string `dbq:"created_at"`
	UpdatedAt   string `dbq:"updated_at"`
}

// opts defines our dbq operations
var opts = &dbq.Options{ConcreteStruct: productQueryModel{}}

// Client is our connection instance
type Client struct {
	database *sql.DB
}

var dt = time.Now()                       // FIXME: temporary for debugging
var timeExample = dt.Format("01-02-2022") // FIXME: temporary for debugging

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

type Row struct {
	ID          string
	GUID        string
	Name        string
	Price       string
	Description string
	CreatedAt   string
	UpdatedAt   string
}
b
// Insert inserts a new row into our database
// FIXME re-add "data interface{}" in parameter
func (c *Client) Insert(ctx *gin.Context) (interface{}, error) {
	stmt := fmt.Sprintf(
		"INSERT INTO products(name, price, description, created_at, updated_at) VALUES(%s, %s, %s, %s, %s, %s, %s)",
		"id", "guid", "name", "price", "description", "created_at", "updated_at",
	)

	res, err := dbq.Q(ctx, c.database, stmt, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert Result == ", res)
	return nil, nil
}

// UpdateById updates a pre-existing row by ID
func (c *Client) UpdateById(ctx *gin.Context, id, username, password, job string) error {
	return nil
}

// FindById finds a user by their id
func (c *Client) FindById(ctx *gin.Context, id string) (*sql.Row, error) {

	return nil, nil
}

// FindAll returns all rows of users
func (c *Client) FindAll(ctx *gin.Context) (*sql.Rows, error) {
	res, err := dbq.Q(ctx, c.database, "SELECT * FROM products", opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products from database: %w", err)
	}
	fmt.Println("results == ", res)
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
