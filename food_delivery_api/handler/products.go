package handler

import (
	"errors"
	"fmt"
	"food_delivery/server/response"
	"food_delivery/service"
	"food_delivery/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
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

func (h *ProductHandler) GetByCategoryID(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.MustGetIDFromVars(r)
	if !ok {
		response.SendBadRequestError(w, nil)
		return
	}

	products, err := h.ProductServiceI.GetProductsByCategoryID(int64(id))
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, products)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.MustGetIDFromVars(r)
	if !ok {
		response.SendBadRequestError(w, nil)
		return
	}

	product, err := h.ProductServiceI.GetProductByID(int64(id))
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, product)
}

func (h *ProductHandler) GetFilteredProducts(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	orderBy := queryParams.Get("order_by")
	openNowStr := queryParams.Get("open_now")
	sortDirection := queryParams.Get("sort_direction")
	search := queryParams.Get("search")
	minPriceStr := queryParams.Get("min_price")
	maxPriceStr := queryParams.Get("max_price")
	minPrice, err := strconv.ParseInt(minPriceStr, 10, 64)
	if err != nil {
		response.SendBadRequestError(w, errors.New("invalid int value for 'min_price'"))
		return
	}

	maxPrice, err := strconv.ParseInt(maxPriceStr, 10, 64)
	if err != nil {
		response.SendBadRequestError(w, errors.New("invalid int value for 'max_price'"))
		return
	}

	categoryIDsStr, _ := queryParams["category_ids"]
	var categoryIDs = strings.Split(categoryIDsStr[0], ",")
	var categoryIDsInt []int

	for _, idStr := range categoryIDs {
		idStr = strings.TrimSpace(idStr)
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("Error converting %s to int: %v\n", idStr, err)
			continue
		}
		categoryIDsInt = append(categoryIDsInt, idInt)
	}

	openNow, err := strconv.ParseBool(openNowStr)
	if err != nil {
		response.SendBadRequestError(w, errors.New("invalid boolean value for 'open_now"))
		return
	}

	products, err := h.ProductServiceI.GetFilteredProducts(search, orderBy, sortDirection, openNow, categoryIDsInt, int(minPrice), int(maxPrice))
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, products)
}
