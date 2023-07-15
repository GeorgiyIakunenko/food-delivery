package service

import (
	"errors"
	"food_delivery/repository"
	"food_delivery/repository/models"
	"food_delivery/server/request"
	"food_delivery/server/response"
)

type OrderService struct {
	OrderRepositoryI   repository.OrderRepositoryI
	ProductRepositoryI repository.ProductRepositoryI
}

type OrderServiceI interface {
	GetAllUserOrdersByID(ID int64) ([]*response.Order, error)
	GetOrderProductsByID(orderID int64, userID int64) ([]*response.OrderProduct, error)
	HasAccessToOrder(userID int64, orderID int64) bool
	CreateOrder(order *request.OrderRequest, userID int64) error
	CancelOrderByID(orderID int64, userID int64) error
}

func NewOrderService(OrderRepositoryI repository.OrderRepositoryI, ProductRepositoryI repository.ProductRepositoryI) OrderServiceI {
	return &OrderService{
		OrderRepositoryI:   OrderRepositoryI,
		ProductRepositoryI: ProductRepositoryI,
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

func (s *OrderService) GetOrderProductsByID(orderID int64, userID int64) ([]*response.OrderProduct, error) {

	if !s.HasAccessToOrder(userID, orderID) {
		return nil, errors.New("unauthorized access to order products")
	}

	orderProducts, err := s.OrderRepositoryI.GetOrderProductsByID(orderID)
	if err != nil {
		return nil, err
	}

	var responseOrderProducts []*response.OrderProduct

	for _, orderProduct := range orderProducts {

		product, err := s.ProductRepositoryI.GetByID(orderProduct.ProductID)
		if err != nil {
			return nil, err
		}

		responseProduct := &response.Product{
			ID:         product.ID,
			Name:       product.Name,
			CategoryID: product.CategoryID,
			Category: &response.Category{
				ID:          product.Category.ID,
				Name:        product.Category.Name,
				Image:       product.Category.Image,
				Description: product.Category.Description,
			},
			SupplierID: product.SupplierID,
			Supplier: &response.Supplier{
				ID:        product.Supplier.ID,
				Name:      product.Supplier.Name,
				Type:      product.Supplier.Type,
				Image:     product.Supplier.Image,
				Address:   product.Supplier.Address,
				OpenTime:  product.Supplier.OpenTime,
				CloseTime: product.Supplier.CloseTime,
			},
			Image:       product.Image,
			Price:       product.Price,
			Description: product.Description,
		}

		responseOrderProduct := &response.OrderProduct{
			OrderID:   orderProduct.OrderID,
			ProductID: orderProduct.ProductID,
			Product:   responseProduct,
			Quantity:  orderProduct.Quantity,
		}

		responseOrderProducts = append(responseOrderProducts, responseOrderProduct)
	}

	return responseOrderProducts, nil
}

func (s *OrderService) HasAccessToOrder(userID int64, orderID int64) bool {
	order, err := s.OrderRepositoryI.GetOrderByID(orderID)
	if err != nil {
		return false
	}

	if order.CustomerID == userID {
		return true
	}

	return false
}

func (s *OrderService) CreateOrder(order *request.OrderRequest, userID int64) error {

	newOrder := &models.Order{
		CustomerID:    userID,
		TotalPrice:    order.TotalPrice,
		OrderStatus:   "PENDING",
		PaymentMethod: order.PaymentMethod,
		Address:       order.Address,
	}

	orderID, err := s.OrderRepositoryI.CreateOrder(newOrder)
	if err != nil {
		return err
	}

	var orderProducts []*models.OrderProduct

	for _, orderProduct := range order.Products {
		newOrderProduct := &models.OrderProduct{
			OrderID:   orderID,
			ProductID: orderProduct.ProductID,
			Quantity:  orderProduct.Quantity,
		}

		orderProducts = append(orderProducts, newOrderProduct)
	}

	err = s.OrderRepositoryI.InsertProductsIntoOrder(orderProducts)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) CancelOrderByID(orderID int64, userID int64) error {
	if !s.HasAccessToOrder(userID, orderID) {
		return errors.New("unauthorized access to order")
	}

	err := s.OrderRepositoryI.CancelOrderByID(orderID)
	if err != nil {
		return err
	}

	return nil
}
