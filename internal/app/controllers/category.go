package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/purplestore/internal/app/schemas"
)

type CategoryServices interface {
	GetList() ([]schemas.CategoryListResponse, error)
	Create(req schemas.CategoryCreateRequest) error
	Detail(req schemas.CategoryDetailRequest) (schemas.CategoryDetailResponse, error)
}

type CategoryController struct {
	service CategoryServices
}

func NewCategoryController(service CategoryServices) *CategoryController {
	return &CategoryController{service: service}
}

func (c *CategoryController) GetList(ctx *gin.Context) {
	response, err := c.service.GetList()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}
	fmt.Println("data categories")
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (c *CategoryController) Create(ctx *gin.Context) {
	var req schemas.CategoryCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err = c.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed create category"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success create category"})
}

func (c *CategoryController) Detail(ctx *gin.Context) {
	categoryIDstr := ctx.Param("id")
	categoryId, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed get data detail"})
	}

	req := schemas.CategoryDetailRequest{ID: categoryId}
	response, err := c.service.Detail(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed get data detail"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})

}
