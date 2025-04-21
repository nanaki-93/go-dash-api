package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-dash-api/internal/model"
	"go-dash-api/internal/service"
	"net/http"
	"strconv"
)

var entityService *service.EntityService

func InitEntityHandler(es *service.EntityService) {
	entityService = es
}

func GetEntities(c echo.Context) error {
	collection := c.Param("collection")

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

	log.Infof("handler - GetEntities - collection:%s page:%s and limit:%s", collection, page, limit)
	entities, err := entityService.GetEntities(collection, pageInt, limitInt)

	entityReqList := make([]*model.Entity, 0, len(entities))
	for _, entity := range entities {
		entityReqList = append(entityReqList, entity)
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entityReqList)
}

func GetEntity(c echo.Context) error {
	collection := c.Param("collection")
	id := c.Param("entityId")
	log.Infof("handler - GetEntity - collection:%s entityId:%s ", collection, id)
	entity, err := entityService.GetEntity(collection, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, entity)
}

func AddEntity(c echo.Context) error {
	entity := new(model.Entity)
	if err := c.Bind(entity); err != nil {
		return err
	}
	log.Infof("handler - AddEntity - entity:%s ", entity)
	err := entityService.AddEntity(entity)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, entity)
}
func DeleteEntity(c echo.Context) error {
	collection := c.Param("collection")
	id := c.Param("entityId")
	log.Infof("handler - DeleteEntity - collection:%s id:%s", collection, id)
	err := entityService.DeleteEntity(collection, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusNoContent, id)
}
