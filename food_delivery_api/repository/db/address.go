package db

import (
	"database/sql"
	"food_delivery/repository/models"
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

/*func (ar *AddressRepository) Create(address models.Address) {

}*/

func (ar *AddressRepository) AddressExists(address models.Address) (bool, error) {
	query := `SELECT EXISTS(
		SELECT 1
		FROM address
		WHERE city = $1
		AND postal_code = $2
		AND address_line1 = $3
		AND address_line2 = $4
		AND country = $5
	)`

	stmt, err := ar.db.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var exists bool

	err = stmt.QueryRow(address.City, address.PostalCode, address.AddressLine1, address.AddressLine2, address.Country)

	if err != nil {
		return false, err
	}

	return exists, nil
}
