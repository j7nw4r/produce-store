package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/j7nw4r/produce-store/models"
	"github.com/j7nw4r/produce-store/schemas"
	"github.com/j7nw4r/produce-store/services"
	"log/slog"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type HttpController struct {
	produceService *services.ProduceService
}

func NewHttpController(produceService *services.ProduceService) HttpController {
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

	idInt, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("could not convert id into uuid")
		c.AbortWithStatusJSON(http.StatusNotFound, "not found")
		return
	}

	produceEntity, err := hc.produceService.GetProduce(c, idInt)
	if err != nil {
		slog.Error(err.Error())
		switch {
		case errors.Is(err, services.ErrNotFound):
			c.AbortWithStatusJSON(http.StatusNotFound, "produce not found")
		case errors.Is(err, services.ErrBadRequest):
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, "error getting services")
		}
		return
	}

	resp := models.FromProduceSchemaToProduce(*produceEntity)
	c.JSON(http.StatusOK, resp)
}

func (hc HttpController) DeleteProduce(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		slog.Error("id (path param) was empty")
		c.AbortWithStatusJSON(http.StatusBadRequest, "id must not be empty")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("could not convert id into uuid")
		c.AbortWithStatusJSON(http.StatusNotFound, "not found")
		return
	}

	produceEntity, err := hc.produceService.DeleteProduce(c, idInt)
	if err != nil {
		slog.Error(err.Error())
		switch {
		case errors.Is(err, services.ErrNotFound):
			c.AbortWithStatusJSON(http.StatusNotFound, "produce not found")
		case errors.Is(err, services.ErrBadRequest):
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, "error getting services")
		}
		return
	}

	resp := models.FromProduceSchemaToProduce(*produceEntity)
	c.JSON(http.StatusOK, resp)
}

func (hc HttpController) SearchProduce(c *gin.Context) {
	name := c.Query("name")
	code := c.Query("code")

	if name != "" && code != "" {
		slog.Error("both name and code params were selected")
		c.AbortWithStatusJSON(http.StatusBadRequest, "must search by either code or name")
		return
	}

	var produceEntities []schemas.ProduceSchema
	var err error
	switch {
	case name != "":
		produceEntities, err = hc.produceService.SearchProduceByName(c, name)
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, "error searching for services")
			return
		}
	case code != "":
		produceEntities, err = hc.produceService.SearchProduceByCode(c, code)
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, "error searching for services")
			return
		}
	default:
		slog.Error("name and code params were empty")
		c.AbortWithStatusJSON(http.StatusBadRequest, "must search by either code or name")
		return
	}

	responses := models.FromProduceSchemasToProduces(produceEntities)
	c.JSON(http.StatusOK, responses)
}

func (hc HttpController) PostProduce(c *gin.Context) {
	pp := []models.Produce{}
	if err := c.Bind(&pp); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, "could not read read body")
		return
	}

	for _, p := range pp {
		if err := validateProduce(&p); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	pSchemas := models.FromProducesToProduceSchemas(pp)
	retProduceSchemas, err := hc.produceService.StoreProduce(c, pSchemas)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, "could not store produce")
		return
	}

	retP := models.FromProduceSchemasToProduces(retProduceSchemas)
	c.JSON(http.StatusOK, retP)
}

func validateProduce(p *models.Produce) error {
	var result error

	if p.Id != 0 {
		result = multierror.Append(result, errors.New("unexpected id"))
	}

	if len(p.Code) != 19 {
		result = multierror.Append(result, errors.New("code is incorrect length"))
	}

	parts := strings.Split(p.Code, "-")
	if len(parts) != 4 {
		result = multierror.Append(result, errors.New("code is incorrect"))
	} else {
		for _, part := range parts {
			if len(part) != 4 {
				result = multierror.Append(result, errors.New("code is incorrect"))
				break
			}
		}
	}

	remainder := math.Remainder(float64(p.Price), .001)
	roundedRemainder := math.Round(remainder)

	if roundedRemainder != 0.0 {
		retErr := fmt.Errorf("remainder was %f", remainder)
		slog.Error(retErr.Error())
		result = multierror.Append(result, fmt.Errorf("price is incorrect: %w", retErr))
	}

	return result
}
