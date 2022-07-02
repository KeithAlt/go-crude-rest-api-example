package product

import "errors"

// ModelCollection defines a collection of models
type ModelCollection struct {
	Repo []Model
}

// NewModelRepo constructs a new model repo
func NewModelRepo() *ModelCollection {
	return &ModelCollection{}
}

// ModelJSONCollection defines a model collection in JSON form
type ModelJSONCollection struct {
	Repo []ModelJSON
}

// ToJSON returns a collection of models in JSON form
func (c *ModelCollection) ToJSON() (*ModelJSONCollection, error) {
	if len(c.Repo) == 0 {
		return nil, errors.New("cannot convert empty collection into JSON")
	}

	var collection ModelJSONCollection
	for _, val := range c.Repo {
		collection.Repo = append(collection.Repo, *val.ToJSON())
	}
	return &collection, nil
}
