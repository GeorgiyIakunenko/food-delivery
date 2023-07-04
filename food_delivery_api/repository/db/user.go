package db

import (
	"database/sql"
	"fmt"
	"food_delivery/repository/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetAll() ([]*models.User, error) {
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

	var users []*models.User

	for rows.Next() {
		user := new(models.User)
		err := scanUser(rows, user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
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

func scanUser(rows *sql.Rows, user *models.User) error {
	return rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Age, &user.Email, &user.Phone, &user.Password, &user.AddressID, &user.CreatedAt)
}
