package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-dash-api/internal/handler"
	"go-dash-api/internal/repository"
	"go-dash-api/internal/router"
	"go-dash-api/internal/service"
	"go-dash-api/internal/validator"
)

func InitServer() *echo.Echo {
	e := echo.New()
	dbConn := NewFirebaseConnection()

	entityRepo := repository.NewFirebaseEntityRepository(dbConn)
	entityService := service.NewEntityService(entityRepo)
	schemaRepo := repository.NewFirebaseSchemaRepository(dbConn)
	schemaService := service.NewSchemaService(schemaRepo)
	e.Validator = validator.NewJsonValidator(schemaService)

	handler.InitEntityHandler(entityService)
	handler.InitSchemaHandler(schemaService)
	router.InitEntityRouter(e, schemaService)
	router.InitSchemaRouter(e)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}
