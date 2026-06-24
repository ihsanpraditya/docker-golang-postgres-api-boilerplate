package router

import (
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", handler.CreateUser)
		userRoutes.GET("", handler.GetUsers)
		userRoutes.GET("/:id", handler.GetUser)
		userRoutes.PUT("/:id", handler.UpdateUser)
		userRoutes.DELETE("/:id", handler.DeleteUser)
	}
}
