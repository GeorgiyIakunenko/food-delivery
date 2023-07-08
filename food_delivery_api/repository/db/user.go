package db

import (
	"database/sql"
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

/*func (r *UserRepository) GetUser(id int64) (*response.User, error) {
	row := r.db.QueryRow("SELECT * FROM customer WHERE id = $1", id)

	user := &response.User{}
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Age, &user.Email, &user.Phone, &user.Address)
	if err != nil {
		return nil, err
	}

	return user, nil
}
*/
