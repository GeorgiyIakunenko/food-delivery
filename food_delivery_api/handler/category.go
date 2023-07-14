package handler

import (
	"food_delivery/server/response"
	"food_delivery/service"
	"food_delivery/utils"
	"net/http"
)

type CategoryHandler struct {
	CategoryServiceI service.CategoryServiceI
}

func NewCategoryHandler(CategoryServiceI service.CategoryServiceI) *CategoryHandler {
	return &CategoryHandler{
		CategoryServiceI: CategoryServiceI,
	}
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.CategoryServiceI.GetAllCategories()
	if err != nil {
		response.SendServerError(w, err)
		return
	}
	response.SendOK(w, categories)
}

func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.MustGetIDFromVars(r)
	if !ok {
		response.SendBadRequestError(w, nil)
		return
	}

	category, err := h.CategoryServiceI.GetCategoryByID(int64(id))
	if err != nil {
		response.SendServerError(w, err)
		return
	}
	response.SendOK(w, category)
}

func (h *CategoryHandler) GetCategoriesBySupplierID(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.MustGetIDFromVars(r)
	if !ok {
		response.SendBadRequestError(w, nil)
		return
	}

	categories, err := h.CategoryServiceI.GetCategoriesBySupplierID(int64(id))
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, categories)
}

func (h *CategoryHandler) GetCategoryByProductID(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.MustGetIDFromVars(r)
	if !ok {
		response.SendBadRequestError(w, nil)
		return
	}

	categories, err := h.CategoryServiceI.GetCategoryByProductID(int64(id))
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, categories)
}
