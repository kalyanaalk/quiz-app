package handler

import (
	"net/http"

	"quiz-app/modules/score_detail"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ScoreDetailHandler struct {
	usecase score_detail.ScoreDetailUsecase
}

func NewScoreDetailHandler(uc score_detail.ScoreDetailUsecase) *ScoreDetailHandler {
	return &ScoreDetailHandler{uc}
}

func RegisterRoutes(r *gin.Engine, h *ScoreDetailHandler) {
	g := r.Group("/score-details")
	g.GET("/:scoreId", h.GetByScoreID)
}

func (h *ScoreDetailHandler) GetByScoreID(c *gin.Context) {
	scoreID, err := uuid.Parse(c.Param("scoreId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid score id"})
		return
	}

	out, err := h.usecase.GetDetails(c, scoreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, out)
}
