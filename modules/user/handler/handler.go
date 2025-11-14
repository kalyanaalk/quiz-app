package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"quiz-app/helpers"
	"quiz-app/middlewares"
	userModel "quiz-app/modules/user"
)

type UserHandler struct {
	UserUsecase userModel.UserUsecase
}

func NewUserHandler(usecase userModel.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: usecase}
}

func RegisterRoutes(r *gin.Engine, h *UserHandler) {
	userGroup := r.Group("/users")

	userGroup.POST("/register", h.Register)
	userGroup.POST("/login", h.Login)

	protected := userGroup.Group("/")
	protected.Use(middlewares.JWTAuthMiddleware())

	protected.GET("/user", h.GetProfile)
	protected.GET("/:id", h.GetUserByID)
}

func (h *UserHandler) Register(c *gin.Context) {
	var input userModel.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserUsecase.Register(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func (h *UserHandler) Login(c *gin.Context) {
	var input userModel.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserUsecase.Login(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	id, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in context"})
		return
	}

	user, err := h.UserUsecase.GetProfile(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func getUserIDFromContext(c *gin.Context) (uuid.UUID, error) {
	val, exists := c.Get("user_id")

	if !exists {
		return uuid.Nil, errors.New("user ID not found in context")
	}

	userID, ok := val.(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("user ID in context is of unexpected type: %T", val)
	}

	return userID, nil
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")

	userID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := h.UserUsecase.GetProfile(context.Background(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
