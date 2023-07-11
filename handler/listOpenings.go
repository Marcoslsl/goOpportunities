package handler

import (
	"github.com/Marcoslsl/goOpportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List Openings
// @Description List the jobs openings
// @Tags Openings
// @Accept json
// @Produce json
// @Sucess 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, 500, "error listing openings")
		return
	}
	sendSucess(ctx, "list-openings", openings)
}
