package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
	"github.com/j7nw4r/produce-store/controllers"
	db2 "github.com/j7nw4r/produce-store/db"
	"github.com/j7nw4r/produce-store/services"
	"log/slog"
)

func main() {
	db, err := db2.NewDB()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// Deps
	produceService := services.NewProduceService(db)
	httpController := controllers.NewHttpController(produceService)

	r := gin.Default()
	r.POST("/services", httpController.PostProduce)
	r.GET("/services/:id", httpController.GetProduce)
	r.GET("/search", httpController.SearchProduce)
	if err := r.Run("localhost:23234"); err != nil {
		slog.Error("%s", err)
	}
}
