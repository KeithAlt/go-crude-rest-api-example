package models

// Product defines our default model state
type Product struct {
	Name        string  `dbq:"name"`
	Price       float32 `dbq:"price"`
	Description string  `dbq:"description"`
	CreatedAt   string  `dbq:"created_at"`
	UpdatedAt   string  `dbq:"updated_at"`
	GUID        string  `dbq:"guid"`
}

// ToJSON returns a model object in JSON
func (m *Product) ToJSON() *ProductJSON {
	mod := ProductJSON(*m)
	return &mod
}

// ProductJSON defines our model in JSON form
type ProductJSON struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	GUID        string  `json:"guid"`
}

// ToModel returns a model object in model form
func (m *ProductJSON) ToModel() *Product {
	mod := Product(*m)
	return &mod
}
