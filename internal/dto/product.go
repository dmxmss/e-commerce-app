package dto

import (
	"time"
)

type Product struct {
	ID int `json:"id"`
	CreatedAt time.Time `json:"createdTime"`
	Name string `json:"name"`
	Description string `json:"description"`
	Vendor string `json:"vendor"`
	Price int `json:"price"`
	Tags []string `json:"tags"`
}

type CreateProductRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Vendor string `json:"vendor"`
	Price int `json:"price"`
	Tags []string `json:"tags"`
}

type CreateProductResponse = Product

type GetProductRequest struct {
	ID *int `json:"id"`
}

type GetProductResponse = Product

type GetProductsRequest struct {
	Name string `json:"name;omitempty"`
	Vendor string `json:"vendor;omitempty"`
	Tags []string `json:"tags"`
	Price int `json:"price;omitempty"`
}

type GetProductsResponse = []Product

type UpdateProductRequest struct {
	ID int `json:"id"`
	Name string `json:"name;omitempty"`
	Description string `json:"description;omitempty"`
	Vendor string `json:"vendor;omitempty"`
	Tags []string `json:"tags;omitempty"`
	Price int `json:"price;omitempty"`
}

type UpdateProductResponse = Product

type DeleteProductRequest struct {
	ID int `json:"id"`
}
