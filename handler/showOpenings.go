package handler

import (
	"fmt"

	"github.com/Marcoslsl/goOpportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(ctx *gin.Context) {
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
	sendSucess(ctx, "find-opening", opening)
}
