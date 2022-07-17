package websockets

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"guessr.net/pkg/jwt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	PUBLISH     = "publish"
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

type Hub struct {
	Clients       []Client
	Subscriptions []Subscription
}

type Client struct {
	Id         string
	Connection *websocket.Conn
}

type EndMessage struct {
	Status string `json:"status"`
}

type Subscription struct {
	Topic  string
	Client *Client
	UserId uint
	*sync.Mutex
}

type Message struct {
	Action string          `json:"action"`
	Topic  string          `json:"topic"`
	Token  string          `json:"token"`
	Data   json.RawMessage `json:"data"`
}

type ErrorLine struct {
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}

func (h *Hub) AddClient(client Client) *Hub {
	h.Clients = append(h.Clients, client)
	return h
}

func (h *Hub) RemoveClient(client Client) *Hub {
	for index, sub := range h.Subscriptions {
		if client.Id == sub.Client.Id {
			h.Subscriptions = append(h.Subscriptions[:index], h.Subscriptions[index+1:]...)
		}
	}
	for index, c := range h.Clients {
		if c.Id == client.Id {
			h.Clients = append(h.Clients[:index], h.Clients[index+1:]...)
		}
	}
	return h
}

func (h *Hub) GetSubscriptions(topic string, client *Client) []Subscription {
	var subscriptionList []Subscription

	for _, subscription := range h.Subscriptions {
		if client != nil {
			if subscription.Client.Id == client.Id && subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}
		} else {
			if subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}
		}
	}
	return subscriptionList
}

func (h *Hub) Subscribe(client *Client, topic string, token string) *Hub {
	clientSubs := h.GetSubscriptions(topic, client)
	u, err := jwt.GetUserFromToken(token)
	if err != nil {
		log.Println(err)
	}
	if u < 1 {
		log.Fatalln("User needs to be signed in.")
	}

	Consume(context.Background(), u, client)
	log.Printf("User with ID %d subscribed to topic %s.\n", u, topic)
	if len(clientSubs) > 0 {
		return h
	}
	newSubscription := Subscription{
		Topic:  topic,
		Client: client,
		UserId: u,
	}
	h.Subscriptions = append(h.Subscriptions, newSubscription)
	return h
}

func (h *Hub) Publish(topic string, message []byte) {
	subscriptions := h.GetSubscriptions(topic, nil)
	for _, sub := range subscriptions {
		err := sub.Client.Send(message)
		if err != nil {
			log.Panic(err)
		}
	}
}

func (h *Hub) Unsubscribe(client *Client, topic string) *Hub {
	for index, sub := range h.Subscriptions {
		if sub.Client.Id == client.Id && sub.Topic == topic {
			h.Subscriptions = append(h.Subscriptions[:index], h.Subscriptions[index+1:]...)
		}
	}
	return h
}

func (h *Hub) HandleReceiveMessage(client Client, payload []byte) *Hub {
	m := Message{}
	err := json.Unmarshal(payload, &m)
	if err != nil {
		log.Println("This is not correct message payload")
		return h
	}
	switch m.Action {
	case PUBLISH:
		h.Publish(m.Topic, m.Data)
	case SUBSCRIBE:
		h.Subscribe(&client, m.Topic, m.Token)
	case UNSUBSCRIBE:
		h.RemoveClient(client)
		fmt.Println("Client want to unsubscribe the topic", m.Topic, client.Id)
	default:
		break
	}
	return h
}

func (client *Client) Send(message []byte) error {
	return client.Connection.WriteMessage(1, message)

}
