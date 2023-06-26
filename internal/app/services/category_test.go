package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/heriant0/purplestore/internal/app/models"
	"github.com/heriant0/purplestore/internal/app/schemas"
	"github.com/heriant0/purplestore/internal/mocks"
	"github.com/magiconair/properties/assert"
)

func TestCategoryService_Detail(t *testing.T) {

	type TestCase struct {
		Name        string
		Given       int
		Data        models.Category
		ExpectData  int
		expectError error
	}

	cases := []TestCase{
		{
			Name:  "when category exist",
			Given: 1,
			Data: models.Category{
				ID:          1,
				Name:        "category 1",
				Description: "description 1",
			},
			ExpectData:  1,
			expectError: nil,
		},
		{
			Name:        "when category not exist",
			Given:       1,
			Data:        models.Category{},
			ExpectData:  0,
			expectError: nil,
		},
		{
			Name:        "when error in categor repository",
			Given:       1,
			Data:        models.Category{},
			ExpectData:  0,
			expectError: errors.New("error query"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockCategoryRepository := mocks.NewMockCategoryRepository(mockCtrl)
			mockCategoryRepository.
				EXPECT().
				GetById(tc.Given).
				Return(tc.Data, tc.expectError)
			
			// call function
			req := schemas.CategoryDetailRequest{ID: tc.Given}
			categoryService := NewCategoryService(mockCategoryRepository)
			response, err := categoryService.Detail(req)

			assert.Equal(t, tc.expectError, err)
			assert.Equal(t, tc.ExpectData, response.ID)
		})
	}

	// t.Run("when category exist", func(t *testing.T) {
	// 	// define mocking
	// 	data := models.Category{
	// 		ID:          1,
	// 		Name:        "Category 1",
	// 		Description: "Description category 1",
	// 	}
	// 	mockHandler := gomock.NewController(t)
	// 	mockCategoryRepository := mocks.NewMockCategoryRepository(mockHandler)
	// 	mockCategoryRepository.
	// 		EXPECT().
	// 		GetById(1).
	// 		Return(data, nil)

	// 	req := schemas.CategoryDetailRequest{ID: 1}
	// 	categoryService := NewCategoryService(mockCategoryRepository)
	// 	response, err := categoryService.Detail(req)

	// 	assert.Equal(t, nil, err)
	// 	assert.Equal(t, data.ID, response.ID)
	// })

	// t.Run("when category repository error", func(t *testing.T) {
	// 	// define mocking
	// 	data := models.Category{}
	// 	mockHandler := gomock.NewController(t)
	// 	mockCategoryRepository := mocks.NewMockCategoryRepository(mockHandler)
	// 	mockCategoryRepository.
	// 		EXPECT().
	// 		GetById(1).
	// 		Return(data, errors.New("error query"))

	// 	req := schemas.CategoryDetailRequest{ID: 1}
	// 	categoryService := NewCategoryService(mockCategoryRepository)
	// 	_, err := categoryService.Detail(req)

	// 	assert.Equal(t, errors.New("error query"), err)
	// })
}
