package routes

import (
	"github.com/gin-gonic/gin"
	"sneakers-app/internal/handlers/auth"
	"sneakers-app/internal/handlers/user"
)

func InitRoutes(routes *gin.Engine) {
	routes.POST("/login", auth.LoginHandler)
	routes.POST("/register", auth.RegisterHandler)

	userRoutes := routes.Group("/users")
	{
		userRoutes.GET("/:id", auth.RequireAuth, user.GetUserById)
		userRoutes.GET("/", auth.RequireAuth, user.GetUsers)
		userRoutes.POST("/", auth.RequireAuth, user.CreateUser)
		userRoutes.PUT("/:id", auth.RequireAuth, user.UpdateUser)
		userRoutes.DELETE("/:id", auth.RequireAuth, user.DeleteUser)
	}
}
