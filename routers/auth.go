package routers

import (
	"guessr.net/controllers"
	"guessr.net/routers/middleware"

	"github.com/gin-gonic/gin"
)


func Routes(route *gin.Engine) {
  userController := new(controllers.UserController)
  auth := route.Group("/auth")
  {
    auth.POST("/users", userController.Register)
    auth.GET("/users/me", userController.CurrentUser)
  }
  auth.Use(middleware.AuthenticationMiddleware())

  tokens := route.Group("/auth")
  {
    tokens.POST("/token/create/", userController.Login)
  }
}


