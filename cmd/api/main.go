package main

import (
	"log"
	"todo_api/internal/config"
	"todo_api/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// import "fmt"

func main() {

	var cfg *config.Config
	var err error
	cfg, err = config.Load()

	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	defer pool.Close()

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"message":  "Todo Api is running",
			"status":   "success",
			"database": "connected",
		})
	})

	router.Run(":" + cfg.Port)

}
