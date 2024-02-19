package cmd

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/j7nw4r/produce-store/controllers"
	db2 "github.com/j7nw4r/produce-store/db"
	"github.com/j7nw4r/produce-store/docs"
	"github.com/j7nw4r/produce-store/services"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
			db, err := db2.NewExternalDB(cmd.Context(), "test.db")
			if err != nil {
				slog.Error(err.Error())
				return
			}
			defer func(db *sql.DB) {
				err := db.Close()
				if err != nil {
					slog.Error(err.Error())
					return
				}
			}(db)

			tableRows, err := db.Query("SELECT name FROM sqlite_master WHERE type ='table' AND name NOT LIKE 'sqlite_%'")
			if err != nil {
				slog.Error(errors.Wrap(err, "could not list tables").Error())
				return
			}
			for tableRows.Next() {
				var tableName string
				if err := tableRows.Scan(&tableName); err != nil {
					slog.Error(errors.Wrap(err, "could not list tables").Error())
					return
				}
				fmt.Printf("table: %s\n", tableName)
			}

			// Deps
			produceService := services.NewProduceService(db)
			httpController := controllers.NewHttpController(&produceService)

			r := gin.Default()
			docs.SwaggerInfo.BasePath = "/"
			r.POST("/produce", httpController.PostProduce)
			r.GET("/produce", httpController.GetAllProduce)
			r.GET("/produce/:id", httpController.GetProduce)
			r.GET("/search", httpController.SearchProduce)
			r.DELETE("/produce/:id", httpController.DeleteProduce)
			r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
			if err := r.Run(":23234"); err != nil {
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
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
