package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (app *application) healthcheck(c *gin.Context) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

  c.JSON(http.StatusOK, env)
}
