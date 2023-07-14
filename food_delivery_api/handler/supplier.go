package handler

import (
	"food_delivery/server/response"
	"food_delivery/service"
	"food_delivery/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type SupplierHandler struct {
	SupplierHandlerI service.SupplierServiceI
}

func NewSupplierHandler(SupplierHandlerI service.SupplierServiceI) *SupplierHandler {
	return &SupplierHandler{
		SupplierHandlerI: SupplierHandlerI,
	}
}

func (h *SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := h.SupplierHandlerI.GetAllSuppliers()
	if err != nil {
		response.SendServerError(w, err)
		return
	}
	response.SendOK(w, suppliers)
}

func (h *SupplierHandler) GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.MustGetIDFromVars(r)
	if !ok {
		response.SendBadRequestError(w, nil)
		return
	}

	supplier, err := h.SupplierHandlerI.GetSupplierByID(int64(id))
	if err != nil {
		response.SendServerError(w, err)
		return
	}
	response.SendOK(w, supplier)
}

func (h *SupplierHandler) GetSupplierByType(w http.ResponseWriter, r *http.Request) {
	supplierType := mux.Vars(r)["type"]

	suppliers, err := h.SupplierHandlerI.GetSuppliersByType(supplierType)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, suppliers)
}

func (h *SupplierHandler) GetSuppliersByCategoryID(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.MustGetIDFromVars(r)
	if !ok {
		response.SendBadRequestError(w, nil)
		return
	}

	suppliers, err := h.SupplierHandlerI.GetSuppliersByCategoryID(int64(id))
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, suppliers)
}
