package handler

import (
	"github.com/labstack/echo/v4"
	"go-dash-api/internal/model"
	"go-dash-api/internal/service"
	"log"
	"net/http"
	"strconv"
)

var schemaService *service.SchemaService

func InitSchemaHandler(schSer *service.SchemaService) {
	schemaService = schSer
}

func GetSchemas(c echo.Context) error {

	page := c.QueryParam("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	limit := c.QueryParam("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	schemas, err := schemaService.GetSchemas(pageInt, limitInt)

	schemaReqList := make([]*model.Schema, 0, len(schemas))
	for _, entity := range schemas {
		schemaReqList = append(schemaReqList, entity)
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schemaReqList)
}

func GetSchema(c echo.Context) error {
	name := c.Param("schemaName")
	schema, err := schemaService.GetSchema(name)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, schema)
}

func AddSchema(c echo.Context) error {
	schema := new(model.Schema)
	if err := c.Bind(schema); err != nil {
		log.Fatal("handler - Bind - AddSchema:", err)
		return err
	}
	err := schemaService.AddSchema(schema)
	if err != nil {
		log.Fatal("handler - AddSchema:", err)
		return err
	}
	return c.JSON(http.StatusCreated, schema)
}

func DeleteSchema(c echo.Context) error {
	id := c.Param("schemaId")
	err := schemaService.DeleteSchema(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusNoContent, id)
}
