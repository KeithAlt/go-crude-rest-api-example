package postgres

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
func (c *Client) Create(ctx *gin.Context, products ...models.ProductModel) (interface{}, error) {
	// FIXME: p.Price not parsing correctly
	var inserts []interface{}
	for _, prod := range products {
		prod.GUID = uuid2.NewString()
		prod.CreatedAt = util.GetTime()
		prod.UpdatedAt = util.GetTime()
		inserts = append(inserts, dbq.Struct(prod))
	}

	stmt := dbq.INSERTStmt("service", []string{"name", "price", "description", "created_at", "updated_at", "guid"}, len(inserts), 1)
	res, err := dbq.E(ctx, c.database, stmt, &c.options, inserts)
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
func (c *Client) Find(ctx *gin.Context, id string) (*models.ProductModel, error) {
	uuid, err := uuid2.Parse(id) // ensures the id is never malicious
	if err != nil {
		return nil, errors.New("invalid product id parameter provided")
	}
	stmt := fmt.Sprintf("SELECT * FROM service WHERE guid='%s' LIMIT 1;", uuid)
	res, err := dbq.Qs(ctx, c.database, stmt, models.ProductModel{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve product with an id of %s from repo", id)
	}
	return res.([]*models.ProductModel)[0], nil
}

// FindAll returns all rows of users
func (c *Client) FindAll(ctx *gin.Context) (*models.ModelCollection, error) {
	res, err := dbq.Qs(ctx, c.database, "SELECT * FROM service", models.ProductModel{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get all service from repo: %w", err)
	}

	collection := models.ModelCollection{}
	prods, ok := res.([]*models.ProductModel)
	if !ok {
		return nil, util.WrapError(err, util.ErrorStatusUnknown, err.Error(), err)
	}

	for _, item := range prods {
		collection.Repo = append(collection.Repo, *item)
	}

	return &collection, nil
}
