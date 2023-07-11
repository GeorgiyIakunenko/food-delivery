package db

import (
	"database/sql"
	"fmt"
	"food_delivery/server/request"
	"food_delivery/server/response"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll() ([]*response.User, error) {
	query := "SELECT id, first_name, last_name, username, age, email, phone, customer_address FROM customer"

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

	var users []*response.User
	for rows.Next() {
		user := &response.User{}
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Age,
			&user.Email,
			&user.Phone,
			&user.Address,
		)
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

func (r *UserRepository) RegisterUser(u request.RegisterRequest) (*response.User, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	query := `
	SELECT EXISTS (
		SELECT 1
		FROM customer
		WHERE email = $1
	)
`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var exists bool

	err = stmt.QueryRow(u.Email).Scan(&exists)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("user with email %s already exists", u.Email)
	}

	query = `INSERT INTO customer (first_name, last_name, username, age, email, phone, password, customer_address) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	stmt, err = r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var id int64

	err = stmt.QueryRow(u.FirstName, u.LastName, u.Username, u.Age, u.Email, u.Phone, password, u.Address).Scan(&id)
	if err != nil {
		return nil, err
	}

	var user response.User

	query = `SELECT id, first_name, last_name, username, password, age, email, phone, customer_address FROM customer WHERE id = $1`

	stmt, err = r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Age, &user.Email, &user.Phone, &user.Address)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

/*func (r *UserRepository) GetByID(id int64) (*response.User, error) {

}*/
