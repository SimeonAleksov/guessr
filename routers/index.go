package routers

import (
   "net/http"

   "guessr.net/controllers"
    "github.com/gin-gonic/gin"
)


func RegisterRoutes(route *gin.Engine) {
  route.NoRoute(func(ctx *gin.Context) {
    ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route not found!"})
  })

  ic := new(controllers.InternalController)

  route.GET("/healthcheck", ic.Healthcheck)
}
