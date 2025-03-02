package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"robin-task/internal/auth"
	"robin-task/internal/handler"
	"robin-task/internal/middleware"
	"robin-task/internal/service"
	"time"
)

func SetupRouter(router *gin.Engine, authService *service.AuthService, taskService *service.TaskService, commentService *service.CommentService, changelogService *service.ChangeLogService, jwtService *auth.JWTService) {
	// Apply middleware
	router.Use(middleware.CORSMiddleware())

	// Rate limiting middleware
	router.Use(middleware.RateLimiter(viper.GetInt("rateLimit.reqPerMin"), time.Minute)) // 10 requests per minute

	// Auth routes
	authHandler := handler.NewAuthHandler(authService)
	api := router.Group("/api")
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(jwtService))
	{
		// Task routes
		taskHandler := handler.NewTaskHandler(taskService, changelogService)
		tasks := protected.Group("/tasks")
		{
			tasks.GET("", taskHandler.GetAllTasks)
			tasks.GET("/:id", taskHandler.GetTaskByID)
			tasks.POST("", taskHandler.CreateTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
			tasks.PATCH("/:id/archive", taskHandler.ArchiveTask) // Archive task

			// Task comments
			commentHandler := handler.NewCommentHandler(commentService)
			tasks.GET("/:id/comments", commentHandler.GetCommentsByTaskID)
			tasks.POST("/:id/comments", commentHandler.CreateComment)
			tasks.PUT("/:id/comments/:commentID", commentHandler.UpdateComment)
			tasks.DELETE("/:id/comments/:commentID", commentHandler.DeleteComment)
		}

		// Change Log routes
		changelogHandler := handler.NewChangeLogHandler(changelogService)
		changeLogs := protected.Group("/tasks/:id/changelog")
		{
			changeLogs.GET("", changelogHandler.GetChangeLogsByTaskID) // Get change logs for task
		}
	}
}
