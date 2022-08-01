package websockets

import (
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

var h = &Hub{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var ControllerConn *kafka.Conn

func ServeWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := Client{
		Id:         uuid.NewV4().String(),
		Connection: conn,
	}
	h.AddClient(client)
	log.Println("Connecting new client.")
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Something went wrong", err)
			h.RemoveClient(client)
			log.Println("total clients and subscriptions ", len(h.Clients), len(h.Subscriptions), messageType)
			return
		}
		go h.HandleReceiveMessage(client, p)
	}
}
