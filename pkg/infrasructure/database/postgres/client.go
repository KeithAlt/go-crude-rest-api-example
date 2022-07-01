package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models/product"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
	"github.com/rocketlaunchr/dbq/v2"
	"log"
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

// opts defines our dbq operations
var opts = &dbq.Options{ConcreteStruct: product.ModelJSON{}}

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
// FIXME re-add "data interface{}" in parameter
func (c *Client) Insert(ctx *gin.Context, p product.ModelJSON) (interface{}, error) {
	// TODO: add util.GetTime()
	// FIXME: p.Price does not work as intended
	stmt := fmt.Sprintf(
		"INSERT INTO products (name, price, description, created_at, updated_at) VALUES('%s', %f, '%s', '%s', '%s')",
		p.Name, p.Price, p.Description, "2022-10-10", "2022-10-10",
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
func (c *Client) FindAll(ctx *gin.Context) ([]*product.ModelJSON, error) {
	res, err := dbq.Qs(ctx, c.database, "SELECT * FROM products", product.Model{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products from database: %w", err)
	}

	products, ok := res.([]*product.ModelJSON)
	if !ok {
		return nil, errors.New("failed to martial database find query results")
	}

	// FIXME debug check
	for _, v := range products {
		fmt.Println(*v)
	}

	return products, nil
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
