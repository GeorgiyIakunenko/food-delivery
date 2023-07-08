package db

import (
	"database/sql"
	"fmt"
	"food_delivery/repository/models"
	"food_delivery/server/request"
)

type AddressRepository struct {
	db *sql.DB
}

func NewAddressRepository(db *sql.DB) *AddressRepository {
	return &AddressRepository{
		db: db,
	}
}

func (ar *AddressRepository) GetAll() ([]*models.Address, error) {
	query := `SELECT * FROM address`
	rows, err := ar.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []*models.Address

	for rows.Next() {
		address := &models.Address{}
		err := rows.Scan(&address.ID, &address.City, &address.PostalCode, &address.AddressLine1, &address.AddressLine2, &address.Country)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return addresses, nil
}

func (ar *AddressRepository) Create(address request.Address) (int64, error) {
	// Check if the address already exists
	exists, id, err := ar.AddressExists(address)
	if err != nil {
		return 0, err
	}

	if exists {
		return id, nil
	}

	query := `INSERT INTO address (city, postal_code, address_line1, address_line2, country)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id`

	stmt, err := ar.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(address.City, address.PostalCode, address.AddressLine1, address.AddressLine2, address.Country).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ar *AddressRepository) AddressExists(address request.Address) (bool, int64, error) {
	query := `SELECT id FROM address WHERE city = $1 AND postal_code = $2 AND address_line1 = $3 AND address_line2 = $4 AND country = $5`

	stmt, err := ar.db.Prepare(query)
	if err != nil {
		return false, 0, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(address.City, address.PostalCode, address.AddressLine1, address.AddressLine2, address.Country).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, 0, nil
		}
		return false, 0, err
	}

	return true, id, nil
}

func (ar *AddressRepository) GetById(id int64) (*models.Address, error) {
	query := `SELECT * FROM address WHERE id = $1`

	stmt, err := ar.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	address := &models.Address{}
	err = stmt.QueryRow(id).Scan(&address.ID, &address.City, &address.PostalCode, &address.AddressLine1, &address.AddressLine2, &address.Country)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (ar *AddressRepository) DeleteById(id int64) (bool, error) {

	UsageTimes, err := ar.IsAddressUsed(id)
	if err != nil {
		return false, err
	}

	fmt.Println(UsageTimes)

	if UsageTimes < 2 {
		query := `DELETE FROM address WHERE id = $1`

		stmt, err := ar.db.Prepare(query)
		if err != nil {
			return false, err
		}
		defer stmt.Close()

		_, err = stmt.Exec(id)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil

}

/*func (ar *AddressRepository) Update()*/

func (ar *AddressRepository) IsAddressUsed(id int64) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM (
			SELECT 1
			FROM "order" o
			LEFT JOIN customer c ON o.customer_id = c.id
			LEFT JOIN supplier s ON s.address_id = c.address_id
			WHERE o.customer_address_id = $1 OR c.address_id = $1 OR s.address_id = $1
			LIMIT 1
		) AS is_used
	`

	var isUsed int

	err := ar.db.QueryRow(query, id).Scan(&isUsed)
	if err != nil {
		return 0, err
	}

	return isUsed, nil
}

func (ar *AddressRepository) UpdateById(id int64, address request.Address) (bool, error) {
	query := `UPDATE address SET city = $1, postal_code = $2, address_line1 = $3, address_line2 = $4, country = $5 WHERE id = $6`

	stmt, err := ar.db.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(address.City, address.PostalCode, address.AddressLine1, address.AddressLine2, address.Country, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
