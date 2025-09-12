package dto

import (
	"net/url"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GetUsersParams struct {
	IDs []int `query:"ids"`
	Name string `query:"name"`
	Email string `query:"email"`
	Page int `query:"page"`
	PerPage int `query:"perPage"`
	SortField string `query:"sortField"`
	SortOrder string `query:"sortOrder"`
	All url.Values `query:"-"`
}

type GetUsersResponse struct {
	Data []User `json:"data"`
	Total int64 `json:"total"`
}

type GetUserResponse = User
