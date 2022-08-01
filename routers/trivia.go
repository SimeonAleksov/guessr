package routers

import (
	"guessr.net/controllers"
	"guessr.net/routers/middleware"

	"github.com/gin-gonic/gin"
)

func TriviaRoutes(route *gin.Engine) {
	triviaController := new(controllers.TriviaController)
	trivia := route.Group("/trivia")
	{
		trivia.GET("/categories/", triviaController.FetchTrivia)
	}
	trivia.Use(middleware.AuthenticationMiddleware())
}
