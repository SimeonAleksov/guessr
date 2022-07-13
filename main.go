 package main

import (
	"time"

	"guessr.net/pkg/config"
	"guessr.net/pkg/database"
	"guessr.net/pkg/logger"
	"guessr.net/routers"

	"github.com/spf13/viper"
)


func main() {
  viper.SetDefault("SERVER_TIMEZONE", "Eu?")
  loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
  time.Local = loc

  if err := config.SetupConfig(); err != nil {
    logger.Fatalf("config error: %s", err)
  }

  if err := database.SetupConnection(); err != nil {
    logger.Fatalf("config error: %s", err)
  }

  router := routers.SetupRoute()

  logger.Fatalf("%v", router.Run(config.ServerConfig()))
}
