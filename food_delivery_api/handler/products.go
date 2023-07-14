package handler

import (
	"food_delivery/server/response"
	"food_delivery/service"
	"net/http"
)

type ProductHandler struct {
	ProductServiceI service.ProductServiceI
}

func NewProductHandler(ProductServiceI service.ProductServiceI) *ProductHandler {
	return &ProductHandler{
		ProductServiceI: ProductServiceI,
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductServiceI.GetAllProducts()
	if err != nil {
		response.SendServerError(w, err)
		return
	}
	response.SendOK(w, products)
}
