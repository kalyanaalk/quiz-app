package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"quiz-app/modules/answer"
)

type AnswerHandler struct {
	UC answer.AnswerUsecase
}

func NewAnswerHandler(uc answer.AnswerUsecase) *AnswerHandler {
	return &AnswerHandler{UC: uc}
}

func RegisterRoutes(r *gin.Engine, h *AnswerHandler) {
	g := r.Group("/answers")

	g.POST("/", h.Create)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)

	g.GET("/question/:question_id", h.GetByQuestionID)
}

func (h *AnswerHandler) Create(c *gin.Context) {
	var input answer.CreateAnswerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a, err := h.UC.Create(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"answer": a})
}

func (h *AnswerHandler) GetByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	a, err := h.UC.GetByID(context.Background(), id)
	if err != nil || a == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "answer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": a})
}

func (h *AnswerHandler) GetByQuestionID(c *gin.Context) {
	qid, _ := uuid.Parse(c.Param("question_id"))
	list, err := h.UC.GetByQuestionID(context.Background(), qid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answers": list})
}

func (h *AnswerHandler) Update(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var input answer.UpdateAnswerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a, err := h.UC.Update(context.Background(), id, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "answer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": a})
}

func (h *AnswerHandler) Delete(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	err := h.UC.Delete(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "answer deleted"})
}
