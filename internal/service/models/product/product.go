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

// ModelJSON defines our model in JSON form
type ModelJSON struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}
