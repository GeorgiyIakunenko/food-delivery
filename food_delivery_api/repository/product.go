package repository

import (
	"database/sql"
	"food_delivery/server/response"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAll() ([]*response.Product, error) {
	query := `
	SELECT p.id, p.name, p.image, p.price, p.description, s.name , c.name , p.supplier_id, p.category_id
	FROM product p 
	INNER JOIN category c ON p.category_id = c.id
	INNER JOIN supplier s ON p.supplier_id = s.id
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

	var products []*response.Product
	for rows.Next() {
		product := &response.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Image,
			&product.Price,
			&product.Description,
			&product.Supplier,
			&product.Category,
			&product.SupplierID,
			&product.CategoryID,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetByID(id int64) (*response.Product, error) {
	query := `
	SELECT p.id, p.name, p.image, p.price, p.description, s.name , c.name
	FROM product p 
	INNER JOIN category c ON p.category_id = c.id
	INNER JOIN supplier s ON p.supplier_id = s.id
	WHERE p.id = ?
`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	product := &response.Product{}
	err = row.Scan(
		&product.ID,
		&product.Name,
		&product.Image,
		&product.Price,
		&product.Description,
		&product.Supplier,
		&product.Category,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) GetBySupplierID(id int64) ([]*response.Product, error) {
	query := `
	SELECT p.id, p.name, p.image, p.price, p.description, p.supplier_id , p.category_id , s.name , c.name
	FROM product p 
	INNER JOIN category c ON p.category_id = c.id
	INNER JOIN supplier s ON p.supplier_id = s.id
	WHERE p.supplier_id = $1
`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*response.Product
	for rows.Next() {
		product := &response.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Image,
			&product.Price,
			&product.Description,
			&product.SupplierID,
			&product.CategoryID,
			&product.Supplier,
			&product.Category,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetByCategoryID(id int64) ([]*response.Product, error) {
	query := `
	SELECT p.id, p.name, p.image, p.price, p.description, s.name , c.name
	FROM product p 
	INNER JOIN category c ON p.category_id = c.id
	INNER JOIN supplier s ON p.supplier_id = s.id
	WHERE p.category_id = ?
`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*response.Product
	for rows.Next() {
		product := &response.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Image,
			&product.Price,
			&product.Description,
			&product.Supplier,
			&product.Category,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
