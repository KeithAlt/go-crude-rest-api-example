package models

// ModelCollection defines a collection of models
type ModelCollection struct {
	Repo []Product
}

// ToJSON returns a collection of models in JSON form
func (c *ModelCollection) ToJSON() *ModelJSONCollection {
	if len(c.Repo) == 0 {
		return nil
	}
	var collection ModelJSONCollection
	for _, val := range c.Repo {
		collection.Repo = append(collection.Repo, *val.ToJSON())
	}
	return &collection
}

// ModelJSONCollection defines a model collection in JSON form
type ModelJSONCollection struct {
	Repo []ProductJSON
}

// ToModel returns a collection of models in JSON form
func (c *ModelJSONCollection) ToModel() *ModelCollection {
	if len(c.Repo) == 0 {
		return nil
	}
	var collection ModelCollection
	for _, m := range c.Repo {
		collection.Repo = append(collection.Repo, *m.ToModel())
	}
	return &collection
}
