package websockets

import (
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

var h = &Hub{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
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
