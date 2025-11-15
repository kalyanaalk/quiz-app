package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"quiz-app/modules/question"
)

type QuestionHandler struct {
	UC question.QuestionUsecase
}

func NewQuestionHandler(uc question.QuestionUsecase) *QuestionHandler {
	return &QuestionHandler{UC: uc}
}

func RegisterRoutes(r *gin.Engine, h *QuestionHandler) {
	g := r.Group("/questions")

	g.POST("/", h.Create)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)

	g.GET("/quiz/:quiz_id", h.GetByQuizID)
}

func (h *QuestionHandler) Create(c *gin.Context) {
	var input question.CreateQuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	q, err := h.UC.Create(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"question": q})
}

func (h *QuestionHandler) GetByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	q, err := h.UC.GetByID(context.Background(), id)
	if err != nil || q == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"question": q})
}

func (h *QuestionHandler) GetByQuizID(c *gin.Context) {
	quizID, _ := uuid.Parse(c.Param("quiz_id"))
	list, err := h.UC.GetByQuizID(context.Background(), quizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": list})
}

func (h *QuestionHandler) Update(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var input question.UpdateQuestionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	q, err := h.UC.Update(context.Background(), id, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"question": q})
}

func (h *QuestionHandler) Delete(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	err := h.UC.Delete(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "question deleted"})
}
