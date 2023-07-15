package repository

import (
	"database/sql"
	"food_delivery/repository/models"
)

type ProductRepository struct {
	db *sql.DB
}

type ProductRepositoryI interface {
	GetAll() ([]*models.Product, error)
	GetByCategoryIDAndSupplierID(categoryID, supplierID int64) ([]*models.Product, error)
	GetBySupplierID(ID int64) ([]*models.Product, error)
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAll() ([]*models.Product, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.category_id,
			c.name AS category_name,
			c.image AS category_image,
			c.description AS category_description,
			p.supplier_id,
			s.name AS supplier_name,
			s.type AS supplier_type,
			s.image AS supplier_image,
			s.supplier_address,
			s.open_time,
			s.close_time,
			p.image,
			p.price,
			p.description
		FROM
			product p
			JOIN category c ON p.category_id = c.id
			JOIN supplier s ON p.supplier_id = s.id
	`

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

	var products []*models.Product

	for rows.Next() {
		product := &models.Product{}
		category := &models.Category{}
		supplier := &models.Supplier{}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.CategoryID,
			&category.Name,
			&category.Image,
			&category.Description,
			&product.SupplierID,
			&supplier.Name,
			&supplier.Type,
			&supplier.Image,
			&supplier.Address,
			&supplier.OpenTime,
			&supplier.CloseTime,
			&product.Image,
			&product.Price,
			&product.Description,
		)
		if err != nil {
			return nil, err
		}

		product.Category = category
		product.Supplier = supplier

		product.Category.ID = product.CategoryID
		product.Supplier.ID = product.SupplierID

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetByCategoryIDAndSupplierID(categoryID, supplierID int64) ([]*models.Product, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.category_id,
			c.name AS category_name,
			c.image AS category_image,
			c.description AS category_description,
			p.supplier_id,
			s.id AS supplier_id,
			s.name AS supplier_name,
			s.type AS supplier_type,
			s.image AS supplier_image,
			s.supplier_address,
			s.open_time,
			s.close_time,
			p.image,
			p.price,
			p.description
		FROM
			product p
			JOIN category c ON p.category_id = c.id
			JOIN supplier s ON p.supplier_id = s.id
		WHERE
			p.category_id = $1
			AND p.supplier_id = $2
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(categoryID, supplierID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product

	for rows.Next() {
		product := &models.Product{}
		category := &models.Category{}
		supplier := &models.Supplier{}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.CategoryID,
			&category.Name,
			&category.Image,
			&category.Description,
			&product.SupplierID,
			&supplier.ID,
			&supplier.Name,
			&supplier.Type,
			&supplier.Image,
			&supplier.Address,
			&supplier.OpenTime,
			&supplier.CloseTime,
			&product.Image,
			&product.Price,
			&product.Description,
		)
		if err != nil {
			return nil, err
		}

		product.Category = category
		product.Supplier = supplier

		product.Category.ID = product.CategoryID
		product.Supplier.ID = product.SupplierID

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetBySupplierID(ID int64) ([]*models.Product, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.category_id,
			c.name AS category_name,
			c.image AS category_image,
			c.description AS category_description,
			p.supplier_id,
			s.id AS supplier_id,
			s.name AS supplier_name,
			s.type AS supplier_type,
			s.image AS supplier_image,
			s.supplier_address,
			s.open_time,
			s.close_time,
			p.image,
			p.price,
			p.description
		FROM
			product p
			JOIN category c ON p.category_id = c.id
			JOIN supplier s ON p.supplier_id = s.id
		WHERE p.supplier_id = $1
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*models.Product

	for rows.Next() {
		product := &models.Product{}
		category := &models.Category{}
		supplier := &models.Supplier{}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.CategoryID,
			&category.Name,
			&category.Image,
			&category.Description,
			&product.SupplierID,
			&supplier.ID,
			&supplier.Name,
			&supplier.Type,
			&supplier.Image,
			&supplier.Address,
			&supplier.OpenTime,
			&supplier.CloseTime,
			&product.Image,
			&product.Price,
			&product.Description,
		)
		if err != nil {
			return nil, err
		}

		product.Category = category
		product.Supplier = supplier

		product.Category.ID = product.CategoryID
		product.Supplier.ID = product.SupplierID

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
