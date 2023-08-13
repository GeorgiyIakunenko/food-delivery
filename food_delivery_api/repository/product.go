package repository

import (
	"database/sql"
	"food_delivery/repository/models"
	"strconv"
	"time"
)

type ProductRepository struct {
	db *sql.DB
}

type ProductRepositoryI interface {
	GetAll() ([]*models.Product, error)
	GetByCategoryIDAndSupplierID(categoryID, supplierID int64) ([]*models.Product, error)
	GetBySupplierID(ID int64) ([]*models.Product, error)
	GetByCategoryID(ID int64) ([]*models.Product, error)
	GetByID(ID int64) (*models.Product, error)
	GetFilteredProducts(searchLine string, orderBy string, sortDirection string, IsOpenNow bool, categoryIDs []int, minPrice int, maxPrice int) ([]*models.Product, error)
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

func (r *ProductRepository) GetByCategoryID(ID int64) ([]*models.Product, error) {
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
		WHERE p.category_id = $1
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

func (r *ProductRepository) GetByID(ID int64) (*models.Product, error) {
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
		WHERE p.id = $1
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	product := &models.Product{}
	category := &models.Category{}
	supplier := &models.Supplier{}

	err = stmt.QueryRow(ID).Scan(
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

	return product, nil
}

func (r *ProductRepository) GetFilteredProducts(searchLine string, orderBy string, sortDirection string, IsOpenNow bool, categoryIDs []int, minPrice int, maxPrice int) ([]*models.Product, error) {
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
		WHERE 1=1
	`
	var params []interface{}

	if len(categoryIDs) > 0 {
		query += " AND p.category_id IN ("
		for i := range categoryIDs {
			params = append(params, categoryIDs[i])
			query += "$" + strconv.Itoa(len(params))
			if i < len(categoryIDs)-1 {
				query += ","
			}
		}
		query += ")"
	}

	if IsOpenNow {
		currentTime := time.Now().Format("15:05")
		params = append(params, currentTime, currentTime)
		query += " AND s.open_time <= $" + strconv.Itoa(len(params)-1)
		query += " AND s.close_time >= $" + strconv.Itoa(len(params))
	}

	if searchLine != "" {
		params = append(params, "%"+searchLine+"%")
		query += " AND p.name ILIKE $" + strconv.Itoa(len(params))
	}

	if minPrice > 0 {
		params = append(params, minPrice)
		query += " AND p.price >= $" + strconv.Itoa(len(params))
	}

	if maxPrice > 0 {
		params = append(params, maxPrice)
		query += " AND p.price <= $" + strconv.Itoa(len(params))
	}

	orderByClause := ""
	if orderBy == "price" && (sortDirection == "asc" || sortDirection == "desc") {
		orderByClause = " ORDER BY p.price " + sortDirection
	}

	query += orderByClause

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(params...)
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

	return products, nil
}
