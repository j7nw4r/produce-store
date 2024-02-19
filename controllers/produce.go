package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/j7nw4r/produce-store/models"
	"github.com/j7nw4r/produce-store/services"
	"log/slog"
	"net/http"
)

type HttpController struct {
	produceService services.ProduceService
}

func NewHttpController(produceService services.ProduceService) HttpController {
	return HttpController{
		produceService: produceService,
	}
}

func (hc HttpController) GetProduce(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		slog.Error("id (path param) was empty")
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must not be empty")
		return
	}

	eid, err := uuid.Parse(id)
	if err != nil {
		slog.Error("could not convert id into uuid")
		c.AbortWithStatusJSON(http.StatusNotFound, "not found")
		return
	}

	produceEntity, err := hc.produceService.GetProduce(eid)
	if err != nil {
		slog.Error(err.Error())
		err := c.AbortWithError(http.StatusInternalServerError, errors.New("error getting services"))
		if err != nil {
			slog.Error("could not abort with error")
		}
		return
	}

	resp := models.FromProduceSchemaToProduce(*produceEntity)
	c.JSON(http.StatusOK, resp)
}

func (hc HttpController) SearchProduce(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		slog.Error("name (query param) was empty")
		c.AbortWithStatusJSON(http.StatusBadRequest, "name parameter required")
		return
	}

	produceEntities, err := hc.produceService.SearchProduce(name)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error searching for services")
		return
	}

	responses := models.FromProduceSchemasToProduces(produceEntities)
	c.JSON(http.StatusOK, responses)
}

func (hc HttpController) PostProduce(c *gin.Context) {
	p := models.Produce{}
	if err := c.Bind(&p); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, "could not read read body")
		return
	}

	pSchema, err := models.FromProduceToProduceSchema(p)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	retProduceSchema, err := hc.produceService.StoreProduce(*pSchema)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, "could not store produce")
		return
	}

	retP := models.FromProduceSchemaToProduce(*retProduceSchema)
	c.JSON(http.StatusOK, retP)
}
