package repository

import (
	"database/sql"
	"food_delivery/repository/models"
)

type SupplierRepository struct {
	db *sql.DB
}

type SupplierRepositoryI interface {
	GetAll() ([]*models.Supplier, error)
	GetByID(id int64) (*models.Supplier, error)
	GetByType(supplierType string) ([]*models.Supplier, error)
	GetByCategoryID(categoryId int64) ([]*models.Supplier, error)
}

func NewSupplierRepository(db *sql.DB) SupplierRepositoryI {
	return &SupplierRepository{
		db: db,
	}
}

func (r *SupplierRepository) GetAll() ([]*models.Supplier, error) {
	query := "SELECT id, name, type, image, supplier_address, open_time, close_time FROM supplier"

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
	var suppliers []*models.Supplier

	for rows.Next() {
		supplier := models.Supplier{}
		err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.Type, &supplier.Image, &supplier.Address, &supplier.OpenTime, &supplier.CloseTime)
		if err != nil {
			return nil, err
		}

		suppliers = append(suppliers, &supplier)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (r *SupplierRepository) GetByID(id int64) (*models.Supplier, error) {
	query := "SELECT id, name, type, image, supplier_address, open_time, close_time FROM supplier WHERE id = $1"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)

	supplier := &models.Supplier{}

	err = row.Scan(&supplier.ID, &supplier.Name, &supplier.Type, &supplier.Image, &supplier.Address, &supplier.OpenTime, &supplier.CloseTime)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (r *SupplierRepository) GetByType(supplierType string) ([]*models.Supplier, error) {
	query := "SELECT id, name, type, image, supplier_address, open_time, close_time FROM supplier WHERE type = $1"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(supplierType)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var suppliers []*models.Supplier

	for rows.Next() {
		supplier := models.Supplier{}
		err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.Type, &supplier.Image, &supplier.Address, &supplier.OpenTime, &supplier.CloseTime)
		if err != nil {
			return nil, err
		}

		suppliers = append(suppliers, &supplier)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (r *SupplierRepository) GetByCategoryID(categoryId int64) ([]*models.Supplier, error) {
	query := "SELECT DISTINCT s.id, s.name, s.type, s.image, s.supplier_address, s.open_time, s.close_time FROM supplier s JOIN product p ON s.id = p.supplier_id JOIN category c ON p.category_id = c.id WHERE c.id = $1"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(categoryId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var suppliers []*models.Supplier

	for rows.Next() {
		supplier := models.Supplier{}
		err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.Type, &supplier.Image, &supplier.Address, &supplier.OpenTime, &supplier.CloseTime)
		if err != nil {
			return nil, err
		}

		suppliers = append(suppliers, &supplier)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return suppliers, nil

}
