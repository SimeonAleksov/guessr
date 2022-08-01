package controllers

import (
	"github.com/gin-gonic/gin"
	triviaServices "guessr.net/models/trivia"
	"net/http"
)

type TriviaController struct{}

func (tc *TriviaController) FetchTrivia(c *gin.Context) {
	tr := triviaServices.GetAllTrivia()
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": tr})
}
