package repository

import (
	"errors"
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/util"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"github.com/rocketlaunchr/dbq/v2"
)

var fields = []string{"name", "price", "description", "created_at", "updated_at", "guid"}

// Create inserts a new row into our repo
// FIXME re-add "data interface{}" in parameter
func (c *Client) Create(ctx *gin.Context, products *models.ModelCollection) (*models.ModelCollection, error) {
	// FIXME: p.Price not parsing correctly
	var inserts []interface{}

	for i := 0; i < len(products.Repo); i++ {
		products.Repo[i].GUID = uuid2.NewString()
		products.Repo[i].CreatedAt = util.GetTime()
		products.Repo[i].UpdatedAt = util.GetTime()
		inserts = append(inserts, dbq.Struct(products.Repo[i]))
	}

	stmt := dbq.INSERTStmt("products", fields, len(inserts), dbq.PostgreSQL)
	_, err := dbq.E(ctx, c.Database, stmt, &c.Options, inserts)
	if err != nil {
		return nil, internal.WrapError(err, internal.ErrorUnknown, err.Error(), err)
	}

	return products, nil
}

// Update updates a pre-existing row by ID
func (c *Client) Update(ctx *gin.Context, id string, newProduct *models.Product) (*models.Product, error) {
	uuid, err := uuid2.Parse(id)
	if err != nil {
		return nil, errors.New("invalid product id parameter provided")
	}

	findStmt := fmt.Sprintf("SELECT * FROM products WHERE guid='%s' LIMIT 1;", uuid)
	res, err := dbq.Qs(ctx, c.Database, findStmt, models.Product{}, nil)
	if err != nil {
		return nil, internal.WrapError(err, internal.ErrorNotFound, err.Error(), err)
	}

	curProduct := res.([]*models.Product)[0]
	mergedModel := util.MergeProductModels(newProduct, curProduct)
	updateStmt := fmt.Sprintf("UPDATE products SET name = $1, price = $2, description = $3, created_at = $4, updated_at = $5, guid = $6 WHERE guid = '%s'", curProduct.GUID)
	results := dbq.MustQ(ctx, c.Database, updateStmt, &c.Options, dbq.Struct(mergedModel))
	fmt.Println("results == ", results) // DEBUG
	// TODO check if results is an error
	return mergedModel, nil
}

// Find finds a product by their guid
func (c *Client) Find(ctx *gin.Context, id string) (*models.Product, error) {
	uuid, err := uuid2.Parse(id) // ensures our id arg is not malicious
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
		return nil, internal.WrapError(err, internal.ErrorUnknown, err.Error(), err)
	}

	for _, item := range prods {
		collection.Repo = append(collection.Repo, *item)
	}

	return &collection, nil
}

// Delete deletes a product by id
func (c *Client) Delete(ctx *gin.Context, id string) (*models.Product, error) {
	uuid, err := uuid2.Parse(id) // ensures our id arg is not malicious
	if err != nil {
		return nil, errors.New("invalid product id parameter provided")
	}

	findStmt := fmt.Sprintf("SELECT * FROM products WHERE guid='%s' LIMIT 1;", uuid)
	res, err := dbq.Qs(ctx, c.Database, findStmt, models.Product{}, nil)
	if err != nil {
		return nil, internal.WrapError(err, internal.ErrorNotFound, err.Error(), err)
	}

	delStmt := fmt.Sprintf("DELETE FROM products WHERE guid='%s';", uuid)
	_, err = dbq.Qs(ctx, c.Database, delStmt, models.Product{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve product with an id of %s from repo: %w", id, err)
	}

	return res.([]*models.Product)[0], nil
}
