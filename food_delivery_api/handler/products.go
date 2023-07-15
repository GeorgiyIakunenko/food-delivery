package handler

import (
	"food_delivery/server/response"
	"food_delivery/service"
	"food_delivery/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func (h *ProductHandler) GetByCategoryIDAndSupplierID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	categoryID, err := strconv.ParseInt(vars["categoryId"], 10, 64)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	supplierID, err := strconv.ParseInt(vars["supplierId"], 10, 64)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	products, err := h.ProductServiceI.GetProductsByCategoryIDAndSupplierID(categoryID, supplierID)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, products)
}

func (h *ProductHandler) GetBySupplierID(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.MustGetIDFromVars(r)
	if !ok {
		response.SendBadRequestError(w, nil)
		return
	}

	products, err := h.ProductServiceI.GetProductsBySupplierID(int64(id))
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, products)
}
