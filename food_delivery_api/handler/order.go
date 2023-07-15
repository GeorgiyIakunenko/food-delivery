package handler

import (
	"food_delivery/config"
	"food_delivery/server/response"
	"food_delivery/service"
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
