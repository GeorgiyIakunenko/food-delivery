package db

import (
	"database/sql"
	"fmt"
	"food_delivery/repository/models"
	"food_delivery/server/response"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetAll() ([]*response.UserResponse, error) {
	query := `SELECT * FROM customer`

	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*response.UserResponse

	for rows.Next() {
		user := new(models.User)
		err := scanUser(rows, user)
		if err != nil {
			return nil, err
		}

		address, err := NewAddressRepository(ur.db).GetById(user.AddressID)
		if err != nil {
			return nil, err
		}

		userResponse := &response.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Age:       user.Age,
			Email:     user.Email,
			Phone:     user.Phone,
			Address: response.AddressResponse{
				City:         address.City,
				PostalCode:   address.PostalCode,
				AddressLine1: address.AddressLine1,
				AddressLine2: address.AddressLine2,
				Country:      address.Country,
			},
		}

		users = append(users, userResponse)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}

func (ur *UserRepository) GetById(id int64) (*models.User, error) {
	query := `SELECT * FROM customer WHERE id = $1`
	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	user := &models.User{}
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName,
		&user.Username, &user.Age, &user.Email, &user.Phone,
		&user.Password, &user.AddressID, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}

		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) DeleteById(id int64) error {
	query := `DELETE FROM customer WHERE id = $1`

	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Update(user *models.User) error {
	query := `
		UPDATE customer
		SET first_name = $2, last_name = $3, username = $4, age = $5, email = $6, phone = $7, password = $8, address_id = $9
		WHERE id = $1
	`

	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.FirstName, user.LastName, user.Username, user.Age, user.Email, user.Phone, user.Password, user.AddressID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Create(user *models.User) error {
	// Check if the user already exists
	exists, err := ur.UserExists(user.Email)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("user with email '%s' already exists", user.Email)
	}

	query := `INSERT INTO customer (first_name, last_name, username, age, email, phone, password, address_id)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Username, user.Age, user.Email, user.Phone, user.Password, user.AddressID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("failed to insert user")
	}

	return nil
}

/*func (ur *UserRepository) CreateAddress(address *request.Address) (int64, error) {

}*/

func (ur *UserRepository) UserExists(email string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM customer WHERE email = $1)"

	var exists bool
	err := ur.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func scanUser(rows *sql.Rows, user *models.User) error {
	return rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Age, &user.Email, &user.Phone, &user.Password, &user.AddressID, &user.CreatedAt)
}
