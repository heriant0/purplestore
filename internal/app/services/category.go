package services

import (
	"github.com/heriant0/purplestore/internal/app/models"
	"github.com/heriant0/purplestore/internal/app/schemas"
)

type CategoryRepository interface {
	GetList() ([]models.Category, error)
}

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetList() ([]schemas.CategoryListResponse, error) {
	var response []schemas.CategoryListResponse

	data, err := s.repo.GetList()
	if err != nil {
		return response, err
	}

	for _, value := range data {
		var resp schemas.CategoryListResponse
		resp.ID = value.ID
		resp.Name = value.Name
		resp.Description = value.Description
		response = append(response, resp)
	}
	return response, nil
}
