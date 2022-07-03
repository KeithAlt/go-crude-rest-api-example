package product

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal"
)

// ModelCollection defines a collection of models
type ModelCollection struct {
	Repo []Model
}

// ToModel returns a collection of models in JSON form
func (c *ModelJSONCollection) ToModel() (*ModelCollection, error) {
	if len(c.Repo) == 0 {
		return nil, internal.NewError(internal.ErrorStatusInvalidArgument, "cannot convert empty JSON collection", nil)
	}
	var collection ModelCollection
	for _, val := range c.Repo {
		collection.Repo = append(collection.Repo, *val.ToModel())
	}
	return &collection, nil
}

// ModelJSONCollection defines a model collection in JSON form
type ModelJSONCollection struct {
	Repo []ModelJSON
}

// ToJSON returns a collection of models in JSON form
func (c *ModelCollection) ToJSON() (*ModelJSONCollection, error) {
	if len(c.Repo) == 0 {
		return nil, internal.NewError(internal.ErrorStatusInvalidArgument, "cannot convert empty model collection", nil)
	}

	var collection ModelJSONCollection
	for _, val := range c.Repo {
		collection.Repo = append(collection.Repo, *val.ToJSON())
	}
	return &collection, nil
}
