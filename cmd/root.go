package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/j7nw4r/produce-store/controllers"
	db2 "github.com/j7nw4r/produce-store/db"
	"github.com/j7nw4r/produce-store/services"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

var (
	// Used for flags.
	localDB bool

	rootCmd = &cobra.Command{
		Use:   "produce-service",
		Short: "Runs produce-service",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := db2.NewExternalDB("sqlite://test.db")
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
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVar(&localDB, "config", true, "should use a local db instance")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
