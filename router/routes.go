package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jffcm/aluraflix-backend/handler"
	"github.com/jffcm/aluraflix-backend/middlewares"
)

func initializeRoutes(r *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	v1 := r.Group("/api/v1")
	{
		// USERS
		v1.POST("/users", handler.CreateUserHandler)
		v1.POST("/users/login", handler.LoginUserHandler)

		v1.Use(middlewares.AuthMiddleware())
		{
			// VIDEOS
			v1.GET("/videos", handler.ListVideosHandler)
			v1.GET("/videos/:id", handler.ShowVideoHandler)
			v1.GET("/videos/free", handler.FreeVideosHandler)
			v1.POST("/videos", handler.CreateVideoHandler)
			v1.DELETE("/videos/:id", handler.DeleteVideoHandler)
			v1.PUT("/videos", handler.UpdateVideoHandler)

			// CATEGORIES
			v1.GET("/categories", handler.ListCategoriesHandler)
			v1.GET("/categories/:id", handler.ShowCategoryHandler)
			v1.GET("/categories/:id/videos", handler.ListVideosByCategoryHandler)
			v1.POST("/categories", handler.CreateCategoryHandler)
			v1.DELETE("/categories/:id", handler.DeleteCategoryHandler)
			v1.PUT("/categories", handler.UpdateCategoryHandler)
		}
	}
}
