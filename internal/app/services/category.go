package services

import (
	"fmt"

	"github.com/heriant0/purplestore/internal/app/models"
	"github.com/heriant0/purplestore/internal/app/schemas"
	log "github.com/sirupsen/logrus"
)

type CategoryRepository interface {
	GetList() ([]models.Category, error)
	Create(data models.Category) error
	GetById(id int) (models.Category, error)
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

func (s *CategoryService) Create(req schemas.CategoryCreateRequest) error {
	data := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.repo.Create(data)
	if err != nil {
		errMsg := fmt.Errorf("category service - err create : %w", err)
		log.Error(errMsg)
		return err
	}

	return nil
}

func (s *CategoryService) Detail(req schemas.CategoryDetailRequest) (schemas.CategoryDetailResponse, error) {
	var response schemas.CategoryDetailResponse

	data, err := s.repo.GetById(req.ID)
	if err != nil {
		errMsg := fmt.Errorf("category service - err detail : %w", err)
		log.Error(errMsg)
		return response, err
	}

	response.ID = data.ID
	response.Name = data.Name
	response.Description = data.Description

	return response, nil
}
