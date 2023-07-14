package service

import (
	"food_delivery/repository"
	"food_delivery/repository/models"
)

type SupplierService struct {
	supplierRepositoryI repository.SupplierRepositoryI
}

type SupplierServiceI interface {
	GetAllSuppliers() ([]*models.Supplier, error)
	GetSupplierByID(id int64) (*models.Supplier, error)
	GetSuppliersByType(supplierType string) ([]*models.Supplier, error)
	GetSuppliersByCategoryID(categoryID int64) ([]*models.Supplier, error)
}

func NewSupplierService(supplierRepositoryI repository.SupplierRepositoryI) SupplierServiceI {
	return &SupplierService{
		supplierRepositoryI: supplierRepositoryI,
	}
}

func (r *SupplierService) GetAllSuppliers() ([]*models.Supplier, error) {
	suppliers, err := r.supplierRepositoryI.GetAll()
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (r *SupplierService) GetSupplierByID(id int64) (*models.Supplier, error) {
	suppliers, err := r.supplierRepositoryI.GetByID(id)
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (r *SupplierService) GetSuppliersByType(supplierType string) ([]*models.Supplier, error) {
	suppliers, err := r.supplierRepositoryI.GetByType(supplierType)
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (r *SupplierService) GetSuppliersByCategoryID(categoryId int64) ([]*models.Supplier, error) {
	suppliers, err := r.supplierRepositoryI.GetByCategoryID(categoryId)
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}
