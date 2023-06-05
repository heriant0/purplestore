package repositories

import (
	"github.com/heriant0/purplestore/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository (db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetList() ([]models.Category, error) {
	var (
		categories   []models.Category
		sqlStatement = "SELECT id, name, description FROM categories"
	)

	// DB Execution
	rows, err := r.DB.Queryx(sqlStatement)
	if err != nil {
		return categories, err
	}
	for rows.Next() {
		var category models.Category
		rows.StructScan(&category) // nolint:errcheck
		categories = append(categories, category)
	}

	return categories, nil
}
