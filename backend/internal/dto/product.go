package dto

import (
	"time"
	"net/url"
)

type Product struct {
	ID int `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Description string `json:"description"`
	Vendor int `json:"vendor_id"`
	Remaining int `json:"remaining"`
	Price int `json:"price"`
	Category int `json:"category_id"`
}

type CreateProductRequest struct {
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Remaining int `json:"remaining"`
	Price int `json:"price"`
	Category int `json:"category_id"`
}

type CreateProductResponse = Product

type GetProductRequest struct {
	ID *int `json:"id"`
}

type GetProductParams struct {
	ID int `query:"id"`
	IDs []int `query:"ids"`
	Page int `query:"page"`
	PerPage int `query:"perPage"`
	SortField string `query:"sortField"`
	SortOrder string `query:"sortOrder"`
	Target string `query:"target"`
	All url.Values `query:"-"`
}

type GetProductsResponse struct {
	Data []Product `json:"data"`
	Total int64 `json:"total"`
}

type UpdateProductRequest struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Remaining int `json:"remaining,omitempty"`
	CategoryID int `json:"category_id,omitempty"`
	Price int `json:"price,omitempty"`
}

type UpdateProductResponse struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Remaining int `json:"remaining,omitempty"`
	CategoryID int `json:"category_id,omitempty"`
	Price int `json:"price,omitempty"`
}

type DeleteProductRequest struct {
	ID int `json:"id"`
}
