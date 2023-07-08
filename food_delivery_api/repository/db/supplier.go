package db

import (
	"database/sql"
	"food_delivery/server/response"
)

type SupplierRepository struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{
		db: db,
	}
}

func (r *SupplierRepository) GetAll() ([]*response.Supplier, error) {
	query := `
	SELECT s.id, s.name, c.name, s.image, s.supplier_address, s.open_time, s.close_time
	FROM supplier s
	INNER JOIN category c ON s.category_id = c.id
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

	var suppliers []*response.Supplier
	for rows.Next() {
		supplier := &response.Supplier{}
		err := rows.Scan(
			&supplier.ID,
			&supplier.Name,
			&supplier.Category,
			&supplier.Image,
			&supplier.Address,
			&supplier.OpenTime,
			&supplier.CloseTime,
		)
		if err != nil {
			return nil, err
		}

		suppliers = append(suppliers, supplier)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (r *SupplierRepository) GetByID(id int64) (*response.Supplier, error) {
	query := `
	SELECT s.id, s.name, c.name, s.image, s.supplier_address, s.open_time, s.close_time
	FROM supplier s
	INNER JOIN category c ON s.category_id = c.id 
	WHERE s.id = $1
`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	supplier := &response.Supplier{}
	err = stmt.QueryRow(id).Scan(
		&supplier.ID,
		&supplier.Name,
		&supplier.Category,
		&supplier.Image,
		&supplier.Address,
		&supplier.OpenTime,
		&supplier.CloseTime,
	)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (r *SupplierRepository) GetByCategory(category string) ([]*response.Supplier, error) {
	query := `
	SELECT s.id, s.name, c.name, s.image, s.supplier_address, s.open_time, s.close_time
	FROM supplier s
	INNER JOIN category c ON s.category_id = c.id 
	WHERE c.name = $1
`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []*response.Supplier
	for rows.Next() {
		supplier := &response.Supplier{}
		err := rows.Scan(
			&supplier.ID,
			&supplier.Name,
			&supplier.Category,
			&supplier.Image,
			&supplier.Address,
			&supplier.OpenTime,
			&supplier.CloseTime,
		)
		if err != nil {
			return nil, err
		}

		suppliers = append(suppliers, supplier)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return suppliers, nil
}
