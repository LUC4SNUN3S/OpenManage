package handler

import (
	"fmt"
	"go-api/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, paramIsRequired("id", "queryParam").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {

		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening not found"))
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening"))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)

}
