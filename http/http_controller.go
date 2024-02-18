package http

import (
	"github.com/gin-gonic/gin"
	"github.com/j7nw4r/produce-store/produce"
	"net/http"
)

type HttpController struct {
	produceService produce.Service
}

func NewHttpController(produceService produce.Service) HttpController {
	return HttpController{
		produceService: produceService,
	}
}

func (hc HttpController) GetProduce(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + id,
	})
}

func (hc HttpController) SearchProduce(c *gin.Context) {
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
	})
}

func (hc HttpController) PostProduce(c *gin.Context) {
}
