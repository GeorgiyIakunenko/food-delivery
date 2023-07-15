package service

import (
	"food_delivery/repository"
	"food_delivery/server/response"
)

type ProductService struct {
	ProductRepositoryI repository.ProductRepositoryI
}

type ProductServiceI interface {
	GetAllProducts() ([]*response.Product, error)
	GetProductsByCategoryIDAndSupplierID(CategoryID, SupplierID int64) ([]*response.Product, error)
	GetProductsBySupplierID(ID int64) ([]*response.Product, error)
}

func NewProductService(ProductRepositoryI repository.ProductRepositoryI) ProductServiceI {
	return &ProductService{
		ProductRepositoryI: ProductRepositoryI,
	}
}

func (s *ProductService) GetAllProducts() ([]*response.Product, error) {
	products, err := s.ProductRepositoryI.GetAll()
	if err != nil {
		return nil, err
	}

	var responseProducts []*response.Product

	for _, product := range products {
		responseProduct := &response.Product{
			ID:         product.ID,
			Name:       product.Name,
			CategoryID: product.CategoryID,
			Category: &response.Category{
				ID:          product.Category.ID,
				Name:        product.Category.Name,
				Image:       product.Category.Image,
				Description: product.Category.Description,
			},
			SupplierID: product.SupplierID,
			Supplier: &response.Supplier{
				ID:        product.Supplier.ID,
				Name:      product.Supplier.Name,
				Type:      product.Supplier.Type,
				Image:     product.Supplier.Image,
				Address:   product.Supplier.Address,
				OpenTime:  product.Supplier.OpenTime,
				CloseTime: product.Supplier.CloseTime,
			},
			Image:       product.Image,
			Price:       product.Price,
			Description: product.Description,
		}

		responseProducts = append(responseProducts, responseProduct)
	}

	return responseProducts, nil
}

func (s *ProductService) GetProductsByCategoryIDAndSupplierID(CategoryID, SupplierID int64) ([]*response.Product, error) {
	products, err := s.ProductRepositoryI.GetByCategoryIDAndSupplierID(CategoryID, SupplierID)
	if err != nil {
		return nil, err
	}

	var responseProducts []*response.Product

	for _, product := range products {
		responseProduct := &response.Product{
			ID:         product.ID,
			Name:       product.Name,
			CategoryID: product.CategoryID,
			Category: &response.Category{
				ID:          product.Category.ID,
				Name:        product.Category.Name,
				Image:       product.Category.Image,
				Description: product.Category.Description,
			},
			SupplierID: product.SupplierID,
			Supplier: &response.Supplier{
				ID:        product.Supplier.ID,
				Name:      product.Supplier.Name,
				Type:      product.Supplier.Type,
				Image:     product.Supplier.Image,
				Address:   product.Supplier.Address,
				OpenTime:  product.Supplier.OpenTime,
				CloseTime: product.Supplier.CloseTime,
			},
			Image:       product.Image,
			Price:       product.Price,
			Description: product.Description,
		}

		responseProducts = append(responseProducts, responseProduct)
	}

	return responseProducts, nil
}

func (s *ProductService) GetProductsBySupplierID(ID int64) ([]*response.Product, error) {
	products, err := s.ProductRepositoryI.GetBySupplierID(ID)
	if err != nil {
		return nil, err
	}

	var responseProducts []*response.Product

	for _, product := range products {
		responseProduct := &response.Product{
			ID:         product.ID,
			Name:       product.Name,
			CategoryID: product.CategoryID,
			Category: &response.Category{
				ID:          product.Category.ID,
				Name:        product.Category.Name,
				Image:       product.Category.Image,
				Description: product.Category.Description,
			},
			SupplierID: product.SupplierID,
			Supplier: &response.Supplier{
				ID:        product.Supplier.ID,
				Name:      product.Supplier.Name,
				Type:      product.Supplier.Type,
				Image:     product.Supplier.Image,
				Address:   product.Supplier.Address,
				OpenTime:  product.Supplier.OpenTime,
				CloseTime: product.Supplier.CloseTime,
			},
			Image:       product.Image,
			Price:       product.Price,
			Description: product.Description,
		}

		responseProducts = append(responseProducts, responseProduct)
	}

	return responseProducts, nil
}
