package repository

import (
	"errors"
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	util2 "github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/util"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"github.com/rocketlaunchr/dbq/v2"
)

var fields = []string{"name", "price", "description", "created_at", "updated_at", "guid"}

// Create inserts a new row into our repo
// FIXME re-add "data interface{}" in parameter
func (c *Client) Create(ctx *gin.Context, products ...models.Product) (interface{}, error) {
	// FIXME: p.Price not parsing correctly
	var inserts []interface{}
	for _, prod := range products {
		prod.GUID = uuid2.NewString()
		prod.CreatedAt = util2.GetTime()
		prod.UpdatedAt = util2.GetTime()
		inserts = append(inserts, dbq.Struct(prod))
	}

	stmt := dbq.INSERTStmt("products", fields, len(inserts), dbq.PostgreSQL)
	res, err := dbq.E(ctx, c.Database, stmt, &c.Options, inserts)
	if err != nil {
		return nil, util2.WrapError(err, util2.ErrorStatusUnknown, err.Error(), err)
	}
	fmt.Println(res) // FIXME debug
	return nil, nil
}

// Update updates a pre-existing row by ID
func (c *Client) Update(ctx *gin.Context, id string, m []interface{}) (*models.Product, error) {
	_, err := uuid2.Parse(id)
	if err != nil {
		return nil, errors.New("invalid product id parameter provided")
	}

	// TODO: clean prepared statement update

	return nil, nil
}

// Find finds a product by their guid
func (c *Client) Find(ctx *gin.Context, id string) (*models.Product, error) {
	uuid, err := uuid2.Parse(id)
	if err != nil {
		return nil, errors.New("invalid product id parameter provided")
	}
	stmt := fmt.Sprintf("SELECT * FROM products WHERE guid='%s' LIMIT 1;", uuid.String())
	res, err := dbq.Qs(ctx, c.Database, stmt, models.Product{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve product with an id of %s from repo", id)
	}
	return res.([]*models.Product)[0], nil
}

// FindAll returns all rows of users
func (c *Client) FindAll(ctx *gin.Context) (*models.ModelCollection, error) {
	res, err := dbq.Qs(ctx, c.Database, "SELECT * FROM products", models.Product{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products from repo: %w", err)
	}

	collection := models.ModelCollection{}
	prods, ok := res.([]*models.Product)
	if !ok {
		return nil, util2.WrapError(err, util2.ErrorStatusUnknown, err.Error(), err)
	}

	for _, item := range prods {
		collection.Repo = append(collection.Repo, *item)
	}

	return &collection, nil
}

// Delete deletes a product by id
func (c *Client) Delete(ctx *gin.Context, id string) error {
	uuid, err := uuid2.Parse(id)
	if err != nil {
		return errors.New("invalid product id parameter provided")
	}
	stmt := fmt.Sprintf("DELETE FROM products WHERE guid='%s';", uuid)
	_, err = dbq.Qs(ctx, c.Database, stmt, models.Product{}, nil)
	if err != nil {
		return fmt.Errorf("failed to retrieve product with an id of %s from repo: %w", id, err)
	}
	return nil
}
