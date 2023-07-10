package handler

import (
	"fmt"

	"github.com/Marcoslsl/goOpportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, 400, errParamsIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, 404, fmt.Sprintf("Opening with id: %s not found", id))
		return
	}

	// Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, 500, fmt.Sprintf("Error deleting opening with id: %s", id))
		return
	}
	sendSucess(ctx, "delete-opening", opening)
}
