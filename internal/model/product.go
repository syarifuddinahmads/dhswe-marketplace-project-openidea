package model

type Product struct {
	Model

	Name        string `json:"name not null"`
	Stock       int    `json:"stock not null"`
	Description string `json:"description"`
}
