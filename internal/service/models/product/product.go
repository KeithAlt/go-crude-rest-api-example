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
		GUID:        m.GUID,
		Name:        m.Name,
		Price:       m.Price,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

// ModelJSON defines our model in JSON form
type ModelJSON struct {
	GUID        string  `json:"guid"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// ToModel returns a model object in model form
func (m *ModelJSON) ToModel() *Model {
	return &Model{
		GUID:        m.GUID,
		Name:        m.Name,
		Price:       m.Price,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
