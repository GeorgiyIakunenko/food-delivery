package repository

import (
	"database/sql"
	"food_delivery/repository/models"
	"github.com/lib/pq"
)

type OrderRepository struct {
	db *sql.DB
}

type OrderRepositoryI interface {
	GetAllUserOrdersByID(ID int64) ([]*models.Order, error)
	GetOrderProductsByID(ID int64) ([]*models.OrderProduct, error)
	GetOrderByID(ID int64) (*models.Order, error)
	CreateOrder(order *models.Order) (int64, error)
	InsertProductsIntoOrder(orderProducts []*models.OrderProduct) error
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

func (r *OrderRepository) GetOrderProductsByID(ID int64) ([]*models.OrderProduct, error) {
	query := `
		SELECT
			op.order_id,
			p.id AS product_id,
			op.quantity
		FROM
			order_product op
			JOIN product p ON op.product_id = p.id
		WHERE
			op.order_id = $1
	`

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

	var orderProducts []*models.OrderProduct
	for rows.Next() {
		orderProduct := new(models.OrderProduct)
		err := rows.Scan(&orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity)
		if err != nil {
			return nil, err
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	return orderProducts, nil
}

func (r *OrderRepository) GetOrderByID(ID int64) (*models.Order, error) {
	query := `SELECT id, customer_id, total_price, payment_method,order_status, created_at,customer_address  FROM "order" WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(ID)

	order := new(models.Order)
	err = row.Scan(&order.ID, &order.CustomerID, &order.TotalPrice, &order.PaymentMethod, &order.OrderStatus, &order.CreatedAt, &order.Address)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) CreateOrder(order *models.Order) (int64, error) {
	query := `INSERT INTO "order" (customer_id, total_price, payment_method,order_status,customer_address) VALUES ($1, $2, $3, $4,$5) RETURNING id`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var ID int64
	err = stmt.QueryRow(order.CustomerID, order.TotalPrice, order.PaymentMethod, order.OrderStatus, order.Address).Scan(&ID)
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func (r *OrderRepository) InsertProductsIntoOrder(orderProducts []*models.OrderProduct) error {
	query := `
		INSERT INTO order_product (order_id, product_id, quantity)
		SELECT unnest($1::bigint[]), unnest($2::bigint[]), unnest($3::bigint[])
	`

	var orderIDs, productIDs, quantities []int64
	for _, orderProduct := range orderProducts {
		orderIDs = append(orderIDs, orderProduct.OrderID)
		productIDs = append(productIDs, orderProduct.ProductID)
		quantities = append(quantities, orderProduct.Quantity)
	}

	_, err := r.db.Exec(query, pq.Array(orderIDs), pq.Array(productIDs), pq.Array(quantities))
	if err != nil {
		return err
	}

	return nil
}
