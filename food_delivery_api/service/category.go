package service

import (
	"food_delivery/repository"
	"food_delivery/server/response"
)

type CategoryService struct {
	CategoryRepositoryI repository.CategoryRepositoryI
}

type CategoryServiceI interface {
	GetAllCategories() ([]*response.Category, error)
	GetCategoryByID(id int64) (*response.Category, error)
	GetCategoriesBySupplierID(supplierID int64) ([]*response.Category, error)
	GetCategoryByProductID(productID int64) (*response.Category, error)
}

func NewCategoryService(CategoryRepositoryI repository.CategoryRepositoryI) CategoryServiceI {
	return &CategoryService{
		CategoryRepositoryI: CategoryRepositoryI,
	}
}

func (s *CategoryService) GetAllCategories() ([]*response.Category, error) {
	categories, err := s.CategoryRepositoryI.GetAll()
	if err != nil {
		return nil, err
	}

	categoriesResponse := make([]*response.Category, 0)
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, &response.Category{
			ID:          category.ID,
			Name:        category.Name,
			Image:       category.Image,
			Description: category.Description,
		})
	}

	return categoriesResponse, nil
}

func (s *CategoryService) GetCategoryByID(id int64) (*response.Category, error) {
	category, err := s.CategoryRepositoryI.GetByID(id)
	if err != nil {
		return nil, err
	}

	categoryResponse := &response.Category{
		ID:          category.ID,
		Name:        category.Name,
		Image:       category.Image,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (s *CategoryService) GetCategoriesBySupplierID(supplierID int64) ([]*response.Category, error) {
	categories, err := s.CategoryRepositoryI.GetBySupplierID(supplierID)
	if err != nil {
		return nil, err
	}

	categoriesResponse := make([]*response.Category, 0)
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, &response.Category{
			ID:          category.ID,
			Name:        category.Name,
			Image:       category.Image,
			Description: category.Description,
		})
	}

	return categoriesResponse, nil
}

func (s *CategoryService) GetCategoryByProductID(productID int64) (*response.Category, error) {
	category, err := s.CategoryRepositoryI.GetByProductID(productID)
	if err != nil {
		return nil, err
	}

	categoryResponse := &response.Category{
		ID:          category.ID,
		Name:        category.Name,
		Image:       category.Image,
		Description: category.Description,
	}

	return categoryResponse, nil
}
