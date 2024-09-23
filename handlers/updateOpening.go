package handler

import (
	"fmt"
	"go-api/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func applyParametersToUpdate(opening *schemas.Opening, request *UpdateOpeningRequest) {
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
}

func UpdateOpeningHandler(ctx *gin.Context) {

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusNotFound, paramIsRequired("id", "queryParam").Error())
		return
	}

	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error request: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening not found"))
	}

	applyParametersToUpdate(&opening, &request)

	if err := db.Save(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error updating opening"))
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
