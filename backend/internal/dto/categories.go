package dto

import (
	"net/url"
)

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type GetCategoriesParams struct {
	IDs []int `query:"ids"`
	All url.Values `query:"-"`
}

type GetCategoriesResponse = []Category
