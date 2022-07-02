package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models/product"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
	"github.com/rocketlaunchr/dbq/v2"
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
var opts = &dbq.Options{ConcreteStruct: product.Model{}}

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
func (c *Client) Insert(ctx *gin.Context, products ...product.Model) (interface{}, error) {
	// FIXME: p.Price does not work as intended
	var inserts []interface{}
	for _, prod := range products {
		inserts = append(inserts, prod)
		fmt.Println("adding: ", prod) // DEBUG
	}

	fmt.Println("inserts == ", inserts)

	stmt := dbq.INSERTStmt("products", []string{"name", "price", "description", "created_at", "updated_at"}, len(inserts))
	res, err := dbq.E(ctx, c.database, stmt, nil, inserts)
	if err != nil {
		return nil, fmt.Errorf("failed to insert products: %w", err)
	}
	fmt.Println("res == ", res) // DEBUG
	return res, nil
}

// UpdateById updates a pre-existing row by ID
func (c *Client) UpdateById(ctx *gin.Context, id, username, password, job string) error {
	return nil
}

// FindById finds a product by their guid
func (c *Client) FindById(ctx *gin.Context, id string) (*product.Model, error) {
	uuid, err := uuid2.Parse(id) // ensures the id is never malicious
	if err != nil {
		return nil, errors.New("invalid product id parameter provided")
	}
	stmt := fmt.Sprintf("SELECT * FROM products WHERE guid='%s' LIMIT 1;", uuid)
	res, err := dbq.Qs(ctx, c.database, stmt, product.Model{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve product with an id of %s from database", id)
	}
	return res.([]*product.Model)[0], nil
}

// FindAll returns all rows of users
func (c *Client) FindAll(ctx *gin.Context) (*product.ModelCollection, error) {
	res, err := dbq.Qs(ctx, c.database, "SELECT * FROM products", product.Model{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products from database: %w", err)
	}

	collection := product.NewModelRepo()
	products, ok := res.([]*product.Model)
	if !ok {
		return nil, errors.New("failed to retrieve database query results")
	}

	for _, item := range products {
		collection.Repo = append(collection.Repo, *item)
	}

	return collection, nil
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
