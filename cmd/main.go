package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/purplestore/internal/app/controllers"
	"github.com/heriant0/purplestore/internal/app/repositories"
	"github.com/heriant0/purplestore/internal/app/services"
	"github.com/heriant0/purplestore/internal/pkg/config"
	"github.com/heriant0/purplestore/internal/pkg/middlewares"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var cfg config.Config
var dbConn *sqlx.DB
var err error

func init() {
	// load configuration based on app.env
	cfg, err = config.LoadConfig()
	if err != nil {
		panic("failed to load config")
	}

	// Create database connection
	dbConn, err = sqlx.Open(cfg.DatabaseDriver, cfg.DatabaseURL)
	if err != nil {
		errMsg := fmt.Errorf("err database connect: %w", err)
		panic(errMsg)
	}

	err = dbConn.Ping()
	if err != nil {
		errMsg := fmt.Errorf("err database ping: %w", err)
		panic(errMsg)
	}

	// setup logrus logging
	loglLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		errMsg := fmt.Errorf("parse log level : %s", err)
		panic(errMsg)
	}
	log.SetLevel(loglLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	// r := gin.Default()
	r := gin.New()
	r.Use(middlewares.LogMiddleware())

	// init repository
	categoryRepository := repositories.NewCategoryRepository(dbConn)

	// init service
	categoryService := services.NewCategoryService(categoryRepository)

	// init controller
	categoryController := controllers.NewCategoryController(categoryService)

	// routes
	r.GET("categories", categoryController.GetList)

	appPort := fmt.Sprintf(":%s", cfg.AppPort)
	r.Run(appPort) // nolint:errcheck
}
