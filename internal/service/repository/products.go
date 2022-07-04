package repository

import (
	"errors"
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/util"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"github.com/rocketlaunchr/dbq/v2"
)

// Create inserts a new row into our repo
// FIXME re-add "data interface{}" in parameter
func (c *Client) Create(ctx *gin.Context, products ...models.Product) (interface{}, error) {
	// FIXME: p.Price not parsing correctly
	var inserts []interface{}
	for _, prod := range products {
		prod.GUID = uuid2.NewString()
		prod.CreatedAt = util.GetTime()
		prod.UpdatedAt = util.GetTime()
		inserts = append(inserts, dbq.Struct(prod))
	}

	stmt := dbq.INSERTStmt("service", []string{"name", "price", "description", "created_at", "updated_at", "guid"}, len(inserts), 1)
	res, err := dbq.E(ctx, c.Database, stmt, &c.Options, inserts)
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
func (c *Client) Find(ctx *gin.Context, id string) (*models.Product, error) {
	uuid, err := uuid2.Parse(id)
	if err != nil {
		return nil, errors.New("invalid product id parameter provided")
	}
	stmt := fmt.Sprintf("SELECT * FROM products WHERE guid='%s' LIMIT 1;", uuid)
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
		return nil, util.WrapError(err, util.ErrorStatusUnknown, err.Error(), err)
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
