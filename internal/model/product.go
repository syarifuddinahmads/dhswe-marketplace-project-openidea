package model

import (
	"time"
)

type Product struct {
	Product_Id    int        `json:"product_id"`
	Name          string     `json:"name" validate:"required,min=5,max=60"`
	Price         float64    `json:"price" validate:"required,min=0"`
	ImageUrl      string     `json:"image_url" validate:"required,url"`
	Condition     string     `json:"condition" validate:"required,oneof=new second"`
	Tags          []string   `json:"tags" validate:"required,min=0"`
	IsPurchasable bool       `json:"is_purchaseable" validate:"required"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
	DeletedAt     *time.Time `db:"deleted_at"`
}
