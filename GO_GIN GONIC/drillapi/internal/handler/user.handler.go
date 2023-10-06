package handler

import (
	"julianmindria/drill/internal/models"
	"julianmindria/drill/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *UserHandler {
	return &UserHandler{r}
}

func (h *UserHandler) PostUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := h.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}

func (h *UserHandler) Getdata(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		handleError(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	response, err := h.GetUser(&user)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"data":        response,
	})
}

func handleError(ctx *gin.Context, statusCode int, description, message string) {
	ctx.JSON(statusCode, gin.H{
		"status":      statusCode,
		"description": description,
		"message":     message,
	})
}
