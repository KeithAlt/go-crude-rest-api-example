package product

// Model defines our default model state
type Model struct {
	Name        string  `dbq:"name"`
	Price       float32 `dbq:"price"`
	Description string  `dbq:"description"`
	CreatedAt   string  `dbq:"created_at"`
	UpdatedAt   string  `dbq:"updated_at"`
	GUID        string  `dbq:"guid"`
}

// ToJSON returns a model object in JSON
func (m *Model) ToJSON() *ModelJSON {
	mod := ModelJSON(*m)
	return &mod
}

// ModelJSON defines our model in JSON form
type ModelJSON struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	GUID        string  `json:"guid"`
}

// ToModel returns a model object in model form
func (m *ModelJSON) ToModel() *Model {
	mod := Model(*m)
	return &mod
}
