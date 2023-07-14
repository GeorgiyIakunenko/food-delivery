package repository

import (
	"database/sql"
	"food_delivery/repository/models"
)

type CategoryRepository struct {
	db *sql.DB
}

type CategoryRepositoryI interface {
	GetAll() ([]*models.Category, error)
	GetByID(id int64) (*models.Category, error)
	GetBySupplierID(supplierID int64) ([]*models.Category, error)
	GetByProductID(productID int64) (*models.Category, error)
}

func NewCategoryRepository(db *sql.DB) CategoryRepositoryI {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetAll() ([]*models.Category, error) {
	query := "SELECT id, name, image, description FROM category"

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

	var categories []*models.Category

	for rows.Next() {
		category := &models.Category{}
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

func (r *CategoryRepository) GetByID(id int64) (*models.Category, error) {
	query := "SELECT id, name, image, description FROM category WHERE id = $1"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	category := &models.Category{}
	err = stmt.QueryRow(id).Scan(
		&category.ID,
		&category.Name,
		&category.Image,
		&category.Description,
	)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) GetBySupplierID(supplierID int64) ([]*models.Category, error) {
	query := `
		SELECT DISTINCT c.id, c.name, c.image, c.description
		FROM category c
		JOIN product p ON c.id = p.category_id
		JOIN supplier s ON p.supplier_id = s.id
		WHERE s.id = $1
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(supplierID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category

	for rows.Next() {
		category := models.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Image, &category.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) GetByProductID(productID int64) (*models.Category, error) {
	query := `
		SELECT c.id, c.name, c.image, c.description
		FROM category c
		JOIN product p ON c.id = p.category_id
		WHERE p.id = $1
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	category := &models.Category{}
	err = stmt.QueryRow(productID).Scan(&category.ID, &category.Name, &category.Image, &category.Description)
	if err != nil {
		return nil, err
	}

	return category, nil
}
