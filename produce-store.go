package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/produce", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/produce", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/search", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(); err != nil {
		slog.Error("%s", err)
	}
}
