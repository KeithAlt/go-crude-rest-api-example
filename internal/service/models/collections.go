package models

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal"
)

// ModelCollection defines a collection of models
type ModelCollection struct {
	Repo []Product
}

// ToJSON returns a collection of models in JSON form
func (c *ModelCollection) ToJSON() (*ModelJSONCollection, error) {
	if len(c.Repo) == 0 {
		return nil, internal.NewError(internal.ErrorInvalidArgument, "cannot convert empty model collection", nil)
	}
	var collection ModelJSONCollection
	for _, val := range c.Repo {
		collection.Repo = append(collection.Repo, *val.ToJSON())
	}
	return &collection, nil
}

// ModelJSONCollection defines a model collection in JSON form
type ModelJSONCollection struct {
	Repo []ProductJSON
}

// ToModel returns a collection of models in JSON form
func (c *ModelJSONCollection) ToModel() (*ModelCollection, error) {
	if len(c.Repo) == 0 {
		return nil, internal.NewError(internal.ErrorInvalidArgument, "cannot convert empty JSON collection", nil)
	}
	var collection ModelCollection
	for _, m := range c.Repo {
		collection.Repo = append(collection.Repo, *m.ToModel())
	}
	return &collection, nil
}
