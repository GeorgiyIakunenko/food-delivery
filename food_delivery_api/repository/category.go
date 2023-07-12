package repository

import (
	"database/sql"
	"food_delivery/server/response"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetAll() ([]*response.Category, error) {
	query := "SELECT * FROM category"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []*response.Category

	for rows.Next() {
		category := &response.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Image,
			&category.Description,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
