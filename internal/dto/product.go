package dto

import (
	"time"
)

type Product struct {
	ID int `json:"id"`
	CreatedAt time.Time `json:"createdTime"`
	UpdatedAt time.Time `json:"updatedTime"`
	Name string `json:"name"`
	Description string `json:"description"`
	Vendor int `json:"vendor"`
	Remaining int `json:"remaining"`
	Price int `json:"price"`
	Category string `json:"category"`
}

type CreateProductRequest struct {
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Remaining int `json:"remaining"`
	Price int `json:"price"`
	Category string `json:"category"`
}

type CreateProductResponse = Product

type GetProductRequest struct {
	ID *int `json:"id"`
}

type GetProductsBy struct {
	ID *int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Vendor *int `json:"vendor,omitempty"`
	Category string `json:"category,omitempty"`
}

type GetProductsResponse = []Product

type UpdateProductRequest struct {
	ID int `json:"id"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Remaining *int `json:"remaining,omitempty"`
	Category string `json:"category,omitempty"`
	Price *int `json:"price,omitempty"`
}

type UpdateProductResponse = Product

type DeleteProductRequest struct {
	ID int `json:"id"`
}
