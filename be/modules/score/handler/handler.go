package handler

import (
	"fmt"
	"net/http"

	"quiz-app/modules/score"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ScoreHandler struct {
	usecase score.ScoreUsecase
}

func NewScoreHandler(uc score.ScoreUsecase) *ScoreHandler {
	return &ScoreHandler{usecase: uc}
}

func RegisterRoutes(r *gin.Engine, h *ScoreHandler) {
	g := r.Group("/scores")
	g.POST("/start", h.Start)
	g.PUT("/:scoreId/submit", h.Submit)
	g.GET("/:scoreId/detail", h.Detail)
	g.GET("/quiz/:quizId/user/:userId", h.ResultsForQuiz)
}

func (h *ScoreHandler) Start(c *gin.Context) {
	var body struct {
		UserID string `json:"user_id"`
		QuizID string `json:"quiz_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := uuid.Parse(body.UserID)
	quizID, _ := uuid.Parse(body.QuizID)

	s, err := h.usecase.Start(c, userID, quizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

func (h *ScoreHandler) Submit(c *gin.Context) {
	scoreID, _ := uuid.Parse(c.Param("scoreId"))
	var answers []score.AnswerSubmission
	if err := c.ShouldBindJSON(&answers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s, err := h.usecase.Submit(c, scoreID, answers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("anshandler: ", answers)

	c.JSON(http.StatusOK, s)
}

func (h *ScoreHandler) Detail(c *gin.Context) {
	scoreID, _ := uuid.Parse(c.Param("scoreId"))
	resp, err := h.usecase.GetDetail(c, scoreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *ScoreHandler) ResultsForQuiz(c *gin.Context) {
	quizID, _ := uuid.Parse(c.Param("quizId"))
	userID, _ := uuid.Parse(c.Param("userId"))

	resp, err := h.usecase.GetResultsForQuiz(c, quizID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
