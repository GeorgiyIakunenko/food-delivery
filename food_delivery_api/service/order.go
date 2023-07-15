package service

import (
	"food_delivery/repository"
	"food_delivery/server/response"
)

type OrderService struct {
	OrderRepositoryI repository.OrderRepositoryI
}

type OrderServiceI interface {
	GetAllUserOrdersByID(ID int64) ([]*response.Order, error)
}

func NewOrderService(OrderRepositoryI repository.OrderRepositoryI) OrderServiceI {
	return &OrderService{
		OrderRepositoryI: OrderRepositoryI,
	}
}

func (s *OrderService) GetAllUserOrdersByID(ID int64) ([]*response.Order, error) {
	orders, err := s.OrderRepositoryI.GetAllUserOrdersByID(ID)
	if err != nil {
		return nil, err
	}

	var responseOrders []*response.Order

	for _, order := range orders {
		responseOrder := &response.Order{
			ID:            order.ID,
			CustomerID:    order.CustomerID,
			TotalPrice:    order.TotalPrice,
			PaymentMethod: order.PaymentMethod,
			OrderStatus:   order.OrderStatus,
			CreatedAt:     order.CreatedAt,
			Address:       order.Address,
		}

		responseOrders = append(responseOrders, responseOrder)
	}

	return responseOrders, nil
}
