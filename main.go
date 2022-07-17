package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	kfk "guessr.net/pkg/websockets"
	"time"

	"github.com/spf13/viper"
	"guessr.net/pkg/config"
	"guessr.net/pkg/database"
	"guessr.net/pkg/logger"
	"guessr.net/routers"
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
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "message-log", 0)
	if err != nil {
		panic(err)
	}
	go kfk.Produce(context.Background())
	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))
	err = conn.Close()
	if err != nil {
		return
	}
}
