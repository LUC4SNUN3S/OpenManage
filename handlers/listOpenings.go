package handler

import (
	"go-api/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)

}
