package db

import (
	"database/sql"
	"food_delivery/server/response"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetAll() ([]*response.Order, error) {
	return nil, nil
}
