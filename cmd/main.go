package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
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
var enforcer *casbin.Enforcer

func init() {
	// load configuration based on app.env
	cfg, err = config.LoadConfig()
	fmt.Println(cfg.SecretKey, "secret key 1")
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

	// casebin enforcer
	enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", "config/rbac_policy.csv")
	if err != nil {
		errMsg := fmt.Errorf("error enforce casbin: %w", err)
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
	userRepository := repositories.NewUserRepository(dbConn)

	// init service
	categoryService := services.NewCategoryService(categoryRepository)
	userService := services.NewUserService(userRepository, []byte(cfg.SecretKey))
	// init controller
	categoryController := controllers.NewCategoryController(categoryService)
	userController := controllers.NewUserController(userService)

	// category routes
	v1Routes := r.Group("api/v1")
	{
		v1Routes.POST("/auth/register", userController.Register)
		v1Routes.POST("/auth/login", userController.Login)

		v1Routes.GET("categories",
			// userController.Auth,
			categoryController.GetList)
		v1Routes.POST("categories",
			middlewares.AuthorizationMiddleware(enforcer, "ani", "categories", "create"),
			categoryController.Create)
		v1Routes.GET("categories/:id", categoryController.Detail)
	}

	appPort := fmt.Sprintf(":%s", cfg.AppPort)
	err := r.Run(appPort)
	if err != nil {
		log.Panic("cannot start the apps")
	}
	// r.Run(appPort) // nolint:errcheck
}
