package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jffcm/aluraflix-backend/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	openRoutes := map[string]bool{
		"/api/v1/users":       true,
		"/api/v1/users/login": true,
		"/api/v1/videos/free": true,
	}

	return func(ctx *gin.Context) {
		if isOpenRoute(ctx.Request.URL.Path, openRoutes) {
			ctx.Next()
			return
		}

		authHeader := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if err := utils.VerifyToken(tokenString); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}

func isOpenRoute(path string, openRoutes map[string]bool) bool {
	for route := range openRoutes {
		if strings.HasPrefix(path, route) {
			return true
		}
	}
	return false
}
