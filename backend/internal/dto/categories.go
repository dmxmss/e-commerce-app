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
	Page int `query:"page"`
	PerPage int `query:"perPage"`
	SortField string `query:"sortField"`
	SortOrder string `query:"sortOrder"`
	All url.Values `query:"-"`
}

type GetCategoriesResponse struct {
	Data []Category `json:"data"`
	Total int64 `json:"total"`
}
