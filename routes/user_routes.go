package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/niteshKrr/gin-framework/controllers"
	"github.com/niteshKrr/gin-framework/middleware"
)

type User_routes struct {
	user_controller controllers.User_controller
}

func (n *User_routes) Init_user_routes(route *gin.Engine) {
	user := route.Group("/user")
	user.POST("/register", n.user_controller.CreateUser())
	user.POST("/login",controllers.Login)
	user.Use(middlewares.AuthMiddleware())

	user.GET("/all", n.user_controller.GetUsers())
	// user.GET("/:id", n.user_controller.GetUserById())
	// user.DELETE("/:id", n.user_controller.DeleteUserById())
	user.PUT("/:id", n.user_controller.UpdateUser())
}
