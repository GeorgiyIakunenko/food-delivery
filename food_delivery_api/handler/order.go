package handler

import (
	"errors"
	"food_delivery/config"
	"food_delivery/server/response"
	"food_delivery/service"
	"food_delivery/utils"
	"net/http"
)

type OrderHandler struct {
	OrderServiceI service.OrderServiceI
}

func NewOrderHandler(OrderServiceI service.OrderServiceI) *OrderHandler {
	return &OrderHandler{
		OrderServiceI: OrderServiceI,
	}
}

func (h *OrderHandler) GetAllUserOrdersByID(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(config.NewConfig().AccessSecret).(*service.JwtCustomClaims)

	orders, err := h.OrderServiceI.GetAllUserOrdersByID(int64(claims.ID))
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, orders)
}

func (h *OrderHandler) GetOrderProductsByID(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(config.NewConfig().AccessSecret).(*service.JwtCustomClaims)

	orderID, ok := utils.MustGetIDFromVars(r)
	if !ok {
		response.SendBadRequestError(w, errors.New("invalid order id"))
		return
	}

	orderProducts, err := h.OrderServiceI.GetOrderProductsByID(int64(orderID), int64(claims.ID))
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, orderProducts)
}
