package product

// Model defines our default model state
type Model struct {
	ID          int     `dbq:"id"`
	GUID        string  `dbq:"guid"`
	Name        string  `dbq:"name"`
	Price       float32 `dbq:"price"`
	Description string  `dbq:"description"`
	CreatedAt   string  `dbq:"created_at"`
	UpdatedAt   string  `dbq:"updated_at"`
}

// ToJSON returns a model object in JSON
func (m *Model) ToJSON() *ModelJSON {
	return &ModelJSON{
		Name:        m.Name,
		Price:       m.Price,
		Description: m.Description,
	}
}

// ModelCollection defines a collection of models
type ModelCollection struct {
	Content []Model
}

// ToJSON returns a collection of models in JSON form
func (c *ModelCollection) ToJSON() *ModelJSONCollection {
	var jsonCol ModelJSONCollection
	for _, val := range c.Content {
		jsonCol.collection = append(jsonCol.collection, *val.ToJSON())
	}
	return &jsonCol
}

// ModelJSON defines our model in JSON form
type ModelJSON struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

// ModelJSONCollection defines a model collection in JSON form
type ModelJSONCollection struct {
	collection []ModelJSON
}
