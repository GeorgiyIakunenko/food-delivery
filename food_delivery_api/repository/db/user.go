package db

import (
	"database/sql"
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
	rows, err := ur.db.Query(query)

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

func scanUser(rows *sql.Rows, user *models.User) error {
	return rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Age, &user.Email, &user.Phone, &user.Password, &user.AddressID, &user.CreatedAt)
}
