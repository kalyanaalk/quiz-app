package handler

import (
	"net/http"

	"quiz-app/modules/quiz"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type QuizHandler struct {
	usecase quiz.QuizUsecase
}

func NewQuizHandler(uc quiz.QuizUsecase) *QuizHandler {
	return &QuizHandler{usecase: uc}
}

func RegisterRoutes(r *gin.Engine, h *QuizHandler) {
	g := r.Group("/quizzes")
	g.POST("/", h.Create)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
	g.GET("/:id", h.GetByID)
	g.GET("/", h.GetAll)
}

func (h *QuizHandler) Create(c *gin.Context) {
	var input quiz.Quiz
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	q, err := h.usecase.Create(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, q)
}

func (h *QuizHandler) Update(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var input quiz.Quiz
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	q, err := h.usecase.Update(c, id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, q)
}

func (h *QuizHandler) Delete(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	err := h.usecase.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func (h *QuizHandler) GetByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	q, err := h.usecase.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, q)
}

func (h *QuizHandler) GetAll(c *gin.Context) {
	list, err := h.usecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}
