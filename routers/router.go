package routers

import (
	"guessr.net/routers/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


func SetupRoute() *gin.Engine {
	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
      gin.SetMode(gin.ReleaseMode)
	}

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())
	// router.Use(middleware.AuthenticationMiddleware())

	RegisterRoutes(router)

	return router
}
