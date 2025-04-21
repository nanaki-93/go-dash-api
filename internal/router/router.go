package router

import (
	"github.com/labstack/echo/v4"
	"go-dash-api/internal/handler"
	middle "go-dash-api/internal/middleware"
	"go-dash-api/internal/service"
	"go-dash-api/internal/validator"
)

func InitEntityRouter(e *echo.Echo, schemaService *service.SchemaService) *echo.Group {

	e.Validator = validator.NewJsonValidator(schemaService)
	g := e.Group("/entity", middle.Validation)
	g.GET("/:collection", handler.GetEntities)
	g.GET("/:collection/:entityId", handler.GetEntity)
	g.POST("/", handler.AddEntity)
	g.DELETE("/:collection/:entityId", handler.DeleteEntity)
	return g
}

func InitSchemaRouter(e *echo.Echo) *echo.Group {
	g := e.Group("/schema")
	g.GET("/", handler.GetSchemas)
	g.GET("/:schemaName", handler.GetSchema)
	g.POST("/", handler.AddSchema)
	g.DELETE("/:schemaId", handler.DeleteSchema)
	return g
}
