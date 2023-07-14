package service

import (
	"food_delivery/repository"
	"food_delivery/server/response"
)

type SupplierService struct {
	supplierRepositoryI repository.SupplierRepositoryI
}

type SupplierServiceI interface {
	GetAllSuppliers() ([]*response.Supplier, error)
	GetSupplierByID(id int64) (*response.Supplier, error)
	GetSuppliersByType(supplierType string) ([]*response.Supplier, error)
	GetSuppliersByCategoryID(categoryID int64) ([]*response.Supplier, error)
}

func NewSupplierService(supplierRepositoryI repository.SupplierRepositoryI) SupplierServiceI {
	return &SupplierService{
		supplierRepositoryI: supplierRepositoryI,
	}
}

func (r *SupplierService) GetAllSuppliers() ([]*response.Supplier, error) {
	suppliers, err := r.supplierRepositoryI.GetAll()
	if err != nil {
		return nil, err
	}

	suppliersResponse := make([]*response.Supplier, 0)

	for _, supplier := range suppliers {
		suppliersResponse = append(suppliersResponse, &response.Supplier{
			ID:        supplier.ID,
			Name:      supplier.Name,
			Type:      supplier.Type,
			Image:     supplier.Image,
			Address:   supplier.Address,
			OpenTime:  supplier.OpenTime,
			CloseTime: supplier.CloseTime,
		})
	}

	return suppliersResponse, nil
}

func (r *SupplierService) GetSupplierByID(id int64) (*response.Supplier, error) {
	supplier, err := r.supplierRepositoryI.GetByID(id)
	if err != nil {
		return nil, err
	}

	supplierResponse := &response.Supplier{
		ID:        supplier.ID,
		Name:      supplier.Name,
		Type:      supplier.Type,
		Image:     supplier.Image,
		Address:   supplier.Address,
		OpenTime:  supplier.OpenTime,
		CloseTime: supplier.CloseTime,
	}

	return supplierResponse, nil
}

func (r *SupplierService) GetSuppliersByType(supplierType string) ([]*response.Supplier, error) {
	suppliers, err := r.supplierRepositoryI.GetByType(supplierType)
	if err != nil {
		return nil, err
	}

	suppliersResponse := make([]*response.Supplier, 0)

	for _, supplier := range suppliers {
		suppliersResponse = append(suppliersResponse, &response.Supplier{
			ID:        supplier.ID,
			Name:      supplier.Name,
			Type:      supplier.Type,
			Image:     supplier.Image,
			Address:   supplier.Address,
			OpenTime:  supplier.OpenTime,
			CloseTime: supplier.CloseTime,
		})
	}

	return suppliersResponse, nil
}

func (r *SupplierService) GetSuppliersByCategoryID(categoryId int64) ([]*response.Supplier, error) {
	suppliers, err := r.supplierRepositoryI.GetByCategoryID(categoryId)
	if err != nil {
		return nil, err
	}

	suppliersResponse := make([]*response.Supplier, 0)

	for _, supplier := range suppliers {
		suppliersResponse = append(suppliersResponse, &response.Supplier{
			ID:        supplier.ID,
			Name:      supplier.Name,
			Type:      supplier.Type,
			Image:     supplier.Image,
			Address:   supplier.Address,
			OpenTime:  supplier.OpenTime,
			CloseTime: supplier.CloseTime,
		})
	}

	return suppliersResponse, nil
}
