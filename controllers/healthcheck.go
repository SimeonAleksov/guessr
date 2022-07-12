package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


type InternalController struct{}


func (ic InternalController) Healthcheck(ctx *gin.Context) { 
	environment := viper.GetBool("DEBUG")
    var currentEnv string
	if environment {
      currentEnv = "development"
	} else {
      currentEnv = "production"
	}
	env := map[string]interface{}{
		"status": "available",
		"system_info": map[string]string{
			"environment": currentEnv,
			"version":     "1.0.0",
		},
	}
    ctx.JSON(http.StatusOK, env)
}
