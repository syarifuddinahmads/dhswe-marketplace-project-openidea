package model

type Product struct {
	Name          string   `json:"name" validate:"required,min=5,max=60"`
	Price         float64  `json:"price" validate:"required,min=0"`
	ImageURL      string   `json:"imageUrl" validate:"required,url"`
	Condition     string   `json:"condition" validate:"required,oneof=new second"`
	Tags          []string `json:"tags" validate:"required,min=0"`
	IsPurchasable bool     `json:"isPurchasable" validate:"required"`
}
