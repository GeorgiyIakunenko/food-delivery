package service

import (
	"food_delivery/repository"
	"food_delivery/server/request"
	"food_delivery/server/response"
)

type UserService struct {
	userRepositoryI repository.UserRepositoryI
}

type UserServiceI interface {
	GetAll() ([]*response.User, error)
	RegisterUser(u request.RegisterRequest) (*response.User, error)
	GetUserByEmail(email string) (*response.User, error)
	GetUserByID(id int) (*response.User, error)
	UpdateUserPasswordById(id int, password string) error
}

func NewUserService(userRepositoryI repository.UserRepositoryI) UserServiceI {
	return &UserService{
		userRepositoryI: userRepositoryI,
	}
}

func (r *UserService) GetAll() ([]*response.User, error) {
	users, err := r.userRepositoryI.GetAll()
	if err != nil {
		return nil, err
	}

	var userResponses []*response.User
	for _, user := range users {
		userResponses = append(userResponses, &response.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Password:  user.Password,
			Age:       user.Age,
			Email:     user.Email,
			Phone:     user.Phone,
			Address:   user.Address,
		})
	}

	return userResponses, nil
}

func (r *UserService) RegisterUser(u request.RegisterRequest) (*response.User, error) {
	user, err := r.userRepositoryI.RegisterUser(u)
	if err != nil {
		return nil, err
	}

	return &response.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Password:  user.Password,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
	}, nil
}

func (r *UserService) GetUserByEmail(email string) (*response.User, error) {
	user, err := r.userRepositoryI.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return &response.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Password:  user.Password,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
	}, nil
}

func (r *UserService) GetUserByID(id int) (*response.User, error) {

	user, err := r.userRepositoryI.GetUserByID(int64(id))
	if err != nil {
		return nil, err
	}

	return &response.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Password:  user.Password,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
	}, nil
}

func (r *UserService) UpdateUserPasswordById(id int, password string) error {
	err := r.userRepositoryI.UpdateUserPasswordByID(int64(id), password)
	if err != nil {
		return err
	}

	return nil

}
