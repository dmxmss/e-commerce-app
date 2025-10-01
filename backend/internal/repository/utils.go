package repository

import (
	e "github.com/dmxmss/e-commerce-app/error"
	"gorm.io/gorm"
	
	"strings"
)

func handlePagination(q *gorm.DB, page, perPage int) {
	if page != 0 && perPage != 0 {
		q = q.Limit(perPage).Offset((page - 1)*perPage)
	}
}

func handleSorting(q *gorm.DB, sortField, sortOrder string, allowedFields []string) error {
	sortField = strings.TrimSpace(strings.ToLower(sortField))
	sortOrder = strings.TrimSpace(strings.ToLower(sortOrder))

	allowed := false
	for _, allowedField := range allowedFields {
		if sortField == allowedField {
			allowed = true
		}
	}
	if !allowed {
		return e.InvalidInputError{Err: "invalid sorting field"}
	}

	if sortOrder != "" && sortOrder != "asc" && sortOrder != "desc" {
		return e.InvalidInputError{Err: "invalid sort order field"}
	}

	if sortOrder == "" {
		sortOrder = "asc" // default value
	}

	if sortField != "" {
		q = q.Order(sortField + " " + sortOrder)
	}

	return nil
}
