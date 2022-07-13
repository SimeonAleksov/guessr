package controllers

import (
	"net/http"

	"guessr.net/models"
    jwt "guessr.net/pkg/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type UserController struct{}


type RegisterInput struct {
  Username string `json:"username" binding:"required"`
  Password string `json:"password" binding:"required"`
}


func verifyPassword(password, hashedPassword string) error {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func (u UserController) Login(c *gin.Context) {
  var input RegisterInput

  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  user_id, err := models.GetUserByUsername(input.Username)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  token, err := jwt.GenerateToken(user_id, false)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  refreshToken, err := jwt.GenerateToken(user_id, true)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "access_token": token,
    "refresh_token": refreshToken,
  })
}


func (u UserController) Register(c *gin.Context) {
  var input RegisterInput

  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  user := models.User{}

  user.Username = input.Username
  user.Password = input.Password

  _, err := user.SaveUser()
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }


  c.JSON(http.StatusOK, gin.H{
    "data": "success",
  })
}


func (u UserController) CurrentUser(c *gin.Context){
	user_id, err := jwt.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
    user, err := models.GetUserByID(user_id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success","data": user})
}
