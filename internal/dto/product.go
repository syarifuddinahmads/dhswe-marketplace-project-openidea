package dto

import "time"

type CreateProductRequest struct {
	Name            string     `json:"name" validate:"required,min=5,max=60"`
	Price           float64    `json:"price" validate:"required,min=0"`
	Image_Url       string     `json:"image_url" validate:"required,url"`
	Condition       string     `json:"condition" validate:"required,oneof=new second"`
	Stock           int        `json:"stock" validate:"required,min=0"`
	Tags            []string   `json:"tags" validate:"required,min=0"`
	Is_Purchaseable bool       `json:"is_purchaseable" validate:"required"`
	Created_At      time.Time  `db:"created_at"`
	Updated_At      *time.Time `db:"updated_at"`
	Deleted_At      *time.Time `db:"deleted_at"`
}

type UpdateProductRequest struct {
	Name            string     `json:"name" validate:"required,min=5,max=60"`
	Price           float64    `json:"price" validate:"required,min=0"`
	Image_Url       string     `json:"image_url" validate:"required,url"`
	Condition       string     `json:"condition" validate:"required,oneof=new second"`
	Stock           int        `json:"stock" validate:"required,min=0"`
	Tags            []string   `json:"tags" validate:"required,min=0"`
	Is_Purchaseable bool       `json:"is_purchaseable" validate:"required"`
	Created_At      time.Time  `db:"created_at"`
	Updated_At      *time.Time `db:"updated_at"`
	Deleted_At      *time.Time `db:"deleted_at"`
}
type CreateProductResponse struct {
	Response string
}
