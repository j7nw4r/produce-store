package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/produce", PostProduce)
	r.GET("/produce/:id", GetProduce)
	r.GET("/search", SearchProduce)
	if err := r.Run(); err != nil {
		slog.Error("%s", err)
	}
}

func GetProduce(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + id,
	})
}

func SearchProduce(c *gin.Context) {
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
	})
}

func PostProduce(c *gin.Context) {
}
