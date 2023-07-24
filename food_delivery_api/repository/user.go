package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/repository/models"
	"food_delivery/server/request"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

type UserRepositoryI interface {
	// is User exist is local function IsUserExists(user request.RegisterRequest) (bool, error)
	GetAll() ([]*models.User, error)
	RegisterUser(u request.RegisterRequest) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int64) (models.User, error)
	UpdateUserPasswordByID(id int64, password string) error
	UpdateUserProfileByID(req models.User, isUserNameChanged bool) error
}

func NewUserRepository(db *sql.DB) UserRepositoryI {
	return &UserRepository{
		db: db,
	}
}

/*func (r *UserRepository) IsUserExists(user request.RegisterRequest) (bool, error) {
	query := "SELECT id FROM customer WHERE email = $1 OR username = $2 OR phone = $3"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Email, user.Username, user.Phone)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}*/

func (r *UserRepository) IsUsernameExists(username string) (bool, error) {
	query := "SELECT id FROM customer WHERE username = $1"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	query := "SELECT id, first_name, last_name, username, age, email, phone, password, customer_address, created_at FROM customer"

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

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Age,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Address,
			&user.CreatedAt,
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

func (r *UserRepository) RegisterUser(u request.RegisterRequest) (*models.User, error) {

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

	var user models.User

	query = `SELECT id, first_name, last_name, username, password, age, email, phone, customer_address, created_at FROM customer WHERE id = $1`

	stmt, err = r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Age, &user.Email, &user.Phone, &user.Address, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, first_name, last_name, username, password, age, email, phone, customer_address, created_at FROM customer WHERE email = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var user models.User

	err = stmt.QueryRow(email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Age, &user.Email, &user.Phone, &user.Address, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(id int64) (models.User, error) {
	query := `SELECT id, first_name, last_name, username, password, age, email, phone, customer_address, created_at FROM customer WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return models.User{}, err
	}

	var user models.User

	err = stmt.QueryRow(id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Age, &user.Email, &user.Phone, &user.Address, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUserPasswordByID(id int64, password string) error {
	query := `UPDATE customer SET password = $1 WHERE id = $2`

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(hashPassword, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUserProfileByID(req models.User, isUserNameChanged bool) error {
	if isUserNameChanged {
		IsUsernameExists, err := r.IsUsernameExists(req.Username)
		if err != nil {
			return err
		}

		if IsUsernameExists {
			return fmt.Errorf("username is already exists")
		}
	}

	query := `UPDATE customer SET first_name = $1, last_name = $2, username = $3, age = $4, email = $5, phone = $6, customer_address = $7 WHERE id = $8`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(req.FirstName, req.LastName, req.Username, req.Age, req.Email, req.Phone, req.Address, req.ID)
	if err != nil {
		return err
	}

	return nil
}
