package routers

import (
   "net/http"

   "github.com/gin-gonic/gin"
)


func RegisterRoutes(route *gin.Engine) {
  route.NoRoute(func(ctx *gin.Context) {
    ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route not found!"})
  })

  InternalRoutes(route)
  Routes(route)
}
