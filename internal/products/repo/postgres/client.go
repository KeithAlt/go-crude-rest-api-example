package postgres

// FIXME: use a faster postgres driver than pq
import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/products"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/products/repo"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/util"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
	"github.com/rocketlaunchr/dbq/v2"
)

// DBConfig is our repo configuration for establishing our connection
type DBConfig struct {
	User    string `binding:"required"`
	Pass    string `binding:"required"`
	Host    string `binding:"required"`
	Port    string `binding:"required"`
	DbType  string `binding:"required"`
	DbName  string `binding:"required"`
	SSLMode string `binding:"required"`
}

// opts defines our dbq option defaults
var opts = &dbq.Options{ConcreteStruct: products.ProductModel{}}

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
		return nil, fmt.Errorf("failed to connect to repo: %w", err)
	}
	return &Client{db}, nil
}

// Create inserts a new row into our repo
// FIXME re-add "data interface{}" in parameter
func (c *Client) Create(ctx *gin.Context, products ...products.ProductModel) (interface{}, error) {
	// FIXME: p.Price not parsing correctly
	var inserts []interface{}
	for _, prod := range products {
		prod.GUID = uuid2.NewString()
		prod.CreatedAt = util.GetTime()
		prod.UpdatedAt = util.GetTime()
		inserts = append(inserts, dbq.Struct(prod))
	}

	stmt := dbq.INSERTStmt("products", []string{"name", "price", "description", "created_at", "updated_at", "guid"}, len(inserts), 1)
	res, err := dbq.E(ctx, c.database, stmt, opts, inserts)
	if err != nil {
		return nil, util.WrapError(err, util.ErrorStatusUnknown, err.Error(), err)
	}
	fmt.Println(res)
	return nil, nil
}

// Update updates a pre-existing row by ID
func (c *Client) Update(ctx *gin.Context, id, username, password, job string) error {
	// TODO implement ...
	return nil
}

// Find finds a product by their guid
func (c *Client) Find(ctx *gin.Context, id string) (*products.ProductModel, error) {
	uuid, err := uuid2.Parse(id) // ensures the id is never malicious
	if err != nil {
		return nil, errors.New("invalid product id parameter provided")
	}
	stmt := fmt.Sprintf("SELECT * FROM products WHERE guid='%s' LIMIT 1;", uuid)
	res, err := dbq.Qs(ctx, c.database, stmt, products.ProductModel{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve product with an id of %s from repo", id)
	}
	return res.([]*products.ProductModel)[0], nil
}

// FindAll returns all rows of users
func (c *Client) FindAll(ctx *gin.Context) (*repo.ModelCollection, error) {
	res, err := dbq.Qs(ctx, c.database, "SELECT * FROM products", products.ProductModel{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products from repo: %w", err)
	}

	collection := repo.ModelCollection{}
	prods, ok := res.([]*products.ProductModel)
	if !ok {
		return nil, util.WrapError(err, util.ErrorStatusUnknown, err.Error(), err)
	}

	for _, item := range prods {
		collection.Repo = append(collection.Repo, *item)
	}

	return &collection, nil
}

// CreateTables will run 'create if not exist' statement migrations with use of dotsql
func (c *Client) CreateTables() error {
	dot, err := dotsql.LoadFromFile("./repo/migrations/create_tables.sql")
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
	dot, err := dotsql.LoadFromFile("./repo/migrations/create_tables.sql")
	if err != nil {
		return util.WrapError(err, util.ErrorStatusNotFound, err.Error(), err)
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