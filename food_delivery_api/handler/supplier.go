package handler

import (
	"database/sql"
)

type SupplierHandler struct {
	db *sql.DB
}

func NewSupplierHandler(db *sql.DB) *SupplierHandler {
	return &SupplierHandler{
		db: db,
	}
}
