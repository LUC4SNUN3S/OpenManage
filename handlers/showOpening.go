package handler

import (
	"fmt"
	"go-api/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningsHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, paramIsRequired("id", "queryParam").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening not found"))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
