package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"quiz-app/config"
	"quiz-app/modules/user"
	user_handler "quiz-app/modules/user/handler"
	user_repository "quiz-app/modules/user/repository"
	user_usecase "quiz-app/modules/user/usecase"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, relying on environment variables")
	}

	config.InitDB()

	err = config.DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	r := gin.Default()

	// frontendURL := os.Getenv("NEXT_PUBLIC_FRONTEND_URL")

	// if frontendURL == "" {
	// 	frontendURL = "http://localhost:3000"
	// 	log.Println("Warning: NEXT_PUBLIC_FRONTEND_URL not set for CORS, defaulting to http://localhost:3000")
	// } else {
	// 	log.Printf("CORS allowed origin set to: %s", frontendURL)
	// }

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	userRepo := user_repository.NewUserRepository(config.DB)
	userUC := user_usecase.NewUserUsecase(userRepo)
	userHandler := user_handler.NewUserHandler(userUC)

	user_handler.RegisterRoutes(r, userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}
