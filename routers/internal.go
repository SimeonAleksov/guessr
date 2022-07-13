package routers

import (
	"guessr.net/controllers"

	"github.com/gin-gonic/gin"
)


func InternalRoutes(route *gin.Engine) {
  ic := new(controllers.InternalController)
  internal := route.Group("/internal")
  {
      internal.GET("/healthcheck", ic.Healthcheck)
  }
}
