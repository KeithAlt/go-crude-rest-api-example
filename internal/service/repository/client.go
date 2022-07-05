package repository

// FIXME: use a faster repository driver than pq
import (
	"database/sql"
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/util"
	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
	"github.com/rocketlaunchr/dbq/v2"
	"log"
)

// Client is our connection instance
type Client struct {
	Database *sql.DB
	Options  dbq.Options
}

// NewClient creates a Client instance for us with values provided through a DBConfig
func NewClient(c config.Database) (*Client, error) {
	psqlInfo := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		c.DbType, c.User, c.Pass, c.Host, c.Port, c.DbName, c.SSLMode,
	)
	db, err := sql.Open(c.DbType, psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to repo: %w", err)
	}
	return &Client{Database: db, Options: dbq.Options{ConcreteStruct: models.Product{}}}, nil
}

// Initialize establishes our repo connection & migrations
func Initialize() *Client {
	cl, err := NewClient(config.DBConfig)
	if err != nil {
		log.Fatal("failed to connect to repo: %w", err) // TODO improve error handling
	}

	// run our harmless migrations
	defer func(client *Client) {
		err := client.CreateTables()
		if err != nil {
			log.Fatal(err) // TODO improve error handling
		}
	}(cl)
	return cl
}

// CreateTables will run 'create if not exist' statement migrations with use of dotsql
func (c *Client) CreateTables() error {
	dot, err := dotsql.LoadFromFile("./database/migrations/create_tables.sql") // TODO replace with config
	if err != nil {
		return fmt.Errorf("failed to load sql scripts: %w", err)
	}

	_, err = dot.Exec(c.Database, "create-uuid-extension")
	if err != nil {
		return fmt.Errorf("failed to execute create extension migration: %w", err)
	}

	_, err = dot.Exec(c.Database, "create-products-table")
	if err != nil {
		return fmt.Errorf("failed to execute create tables migration: %w", err)
	}
	return nil
}

// DropTables will run 'drop if not exist' statement migrations with use of dotsql
func (c *Client) DropTables() error {
	dot, err := dotsql.LoadFromFile("./database/migrations/create_tables.sql") // TODO replace with config
	if err != nil {
		return util.WrapError(err, util.ErrorStatusNotFound, err.Error(), err)
	}

	_, err = dot.Exec(c.Database, "drop-uuid-extension")
	if err != nil {
		return fmt.Errorf("failed to execute drop extension migration: %w", err)
	}

	_, err = dot.Exec(c.Database, "drop-products-table")
	if err != nil {
		return fmt.Errorf("failed to execute drop tables migration: %w", err)
	}
	return nil
}
