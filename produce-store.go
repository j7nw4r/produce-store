package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
	"github.com/j7nw4r/produce-store/http"
	"github.com/j7nw4r/produce-store/produce"
	"log/slog"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	produceService := produce.NewProduceService(db)
	httpController := http.NewHttpController(produceService)

	r := gin.Default()
	r.POST("/produce", httpController.PostProduce)
	r.GET("/produce/:id", httpController.GetProduce)
	r.GET("/search", httpController.SearchProduce)
	if err := r.Run(); err != nil {
		slog.Error("%s", err)
	}
}
