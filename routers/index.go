package routers

import (
    "net/http"

    "guessr.net/pkg/websockets"

    "github.com/gin-gonic/gin"
)


func RegisterRoutes(route *gin.Engine) {
  route.NoRoute(func(ctx *gin.Context) {
    ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route not found!"})
  })
  h := websockets.GetHub()
  InternalRoutes(route)
  Routes(route)

  go h.Run()
  route.GET(
    "/ws/:roomId", func(ctx *gin.Context) {
      roomId := ctx.Param("roomId")
      websockets.ServeWs(ctx.Writer, ctx.Request, roomId)
    },
  )
}
