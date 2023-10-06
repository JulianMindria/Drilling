package routers

import (
	"julianmindria/drill/internal/handler"
	"julianmindria/drill/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func User(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(d)
	handlers := handler.NewUser(repo)

	route.POST("/", handlers.PostUser)
	route.GET("/", handlers.Getdata)
}
