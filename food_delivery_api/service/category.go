package service

import (
	"food_delivery/repository"
	"food_delivery/repository/models"
)

type CategoryService struct {
	CategoryRepositoryI repository.CategoryRepositoryI
}

type CategoryServiceI interface {
	GetAllCategories() ([]*models.Category, error)
	GetCategoryByID(id int64) (*models.Category, error)
	GetCategoriesBySupplierID(supplierID int64) ([]*models.Category, error)
	GetCategoryByProductID(productID int64) (*models.Category, error)
}

func NewCategoryService(CategoryRepositoryI repository.CategoryRepositoryI) CategoryServiceI {
	return &CategoryService{
		CategoryRepositoryI: CategoryRepositoryI,
	}
}

func (s *CategoryService) GetAllCategories() ([]*models.Category, error) {
	categories, err := s.CategoryRepositoryI.GetAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) GetCategoryByID(id int64) (*models.Category, error) {
	category, err := s.CategoryRepositoryI.GetByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) GetCategoriesBySupplierID(supplierID int64) ([]*models.Category, error) {
	categories, err := s.CategoryRepositoryI.GetBySupplierID(supplierID)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) GetCategoryByProductID(productID int64) (*models.Category, error) {
	category, err := s.CategoryRepositoryI.GetByProductID(productID)
	if err != nil {
		return nil, err
	}

	return category, nil
}
