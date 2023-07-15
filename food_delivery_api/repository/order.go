package repository

import (
	"database/sql"
	"food_delivery/repository/models"
)

type OrderRepository struct {
	db *sql.DB
}

type OrderRepositoryI interface {
	GetAllUserOrdersByID(ID int64) ([]*models.Order, error)
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetAllUserOrdersByID(ID int64) ([]*models.Order, error) {
	query := `SELECT id, customer_id, total_price, payment_method,order_status, created_at,customer_address  FROM "order" WHERE customer_id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order := new(models.Order)
		err := rows.Scan(&order.ID, &order.CustomerID, &order.TotalPrice, &order.PaymentMethod, &order.OrderStatus, &order.CreatedAt, &order.Address)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
