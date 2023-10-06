package routers

import (
	"julianmindria/drill/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/products", handler.GetallProduct)
	User(router, db)
	// router.Use(cors.New(config.CorsConfig))
	// router.Use(cors.Default())
	return router
}
