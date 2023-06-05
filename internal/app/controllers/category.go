package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/purplestore/internal/app/schemas"
)

type CategoryServices interface {
	GetList() ([]schemas.CategoryListResponse, error)
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

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
