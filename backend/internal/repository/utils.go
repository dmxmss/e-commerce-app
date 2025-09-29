package repository

import (
	"gorm.io/gorm"
)

func handlePagination(q *gorm.DB, page, perPage int) {
	if page != 0 && perPage != 0 {
		q = q.Limit(perPage).Offset((page - 1)*perPage)
	}
}

func handleSorting(q *gorm.DB, sortField, sortOrder string) {
	if sortField != "" && sortOrder != "" {
		q = q.Order(sortField + " " + sortOrder)
	}
}
