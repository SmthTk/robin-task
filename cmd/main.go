package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"robin-task/internal/auth"
	"robin-task/internal/database"
	"robin-task/internal/repository"
	"robin-task/internal/router"
	"robin-task/internal/service"
)

func initConfig() {
	// Set the configuration file
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	// Initialize Viper
	initConfig()

	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	changelogRepo := repository.NewChangeLogRepository(db)

	// Initialize services
	jwtService := auth.NewJWTService(viper.GetString("jwt.secret"), viper.GetInt("jwt.expHour")) // 24 hours token expiry
	authService := service.NewAuthService(userRepo, jwtService)
	taskService := service.NewTaskService(taskRepo)
	commentService := service.NewCommentService(commentRepo)
	changelogService := service.NewChangeLogService(changelogRepo)

	// Create router
	ginRouter := gin.Default()
	router.SetupRouter(ginRouter, authService, taskService, commentService, changelogService, jwtService)

	// Start server
	srv := &http.Server{
		Addr:    ":" + viper.GetString("server.port"),
		Handler: ginRouter,
	}

	// Run server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
