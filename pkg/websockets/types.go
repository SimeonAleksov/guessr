package websockets

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"
	models "guessr.net/models/trivia"
	userServices "guessr.net/models/users"
	"guessr.net/pkg/jwt"
	"log"
	"sync"
)

const (
	PUBLISH        = "publish"
	SUBSCRIBE      = "subscribe"
	UNSUBSCRIBE    = "unsubscribe"
	CREATE_OR_JOIN = "create_or_join"
	START          = "start"
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
	Topic  string  `json:"-"`
	Client *Client `json:"-"`
	UserId uint
	*sync.Mutex
}

type Message struct {
	Action string          `json:"action"`
	Topic  string          `json:"topic"`
	Token  string          `json:"token"`
	Data   json.RawMessage `json:"data"`
}

type CreateGameMessage struct {
	TriviaID int64  `json:"triviaID"`
	Code     string `json:"code"`
}

type ErrorLine struct {
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}

type KafkaController struct {
	conn *kafka.Conn
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

	if len(clientSubs) > 0 {
		return h
	}
	newSubscription := Subscription{
		Topic:  topic,
		Client: client,
		UserId: uint(u),
	}
	h.Subscriptions = append(h.Subscriptions, newSubscription)
	h.CurrentUsers(topic)
	go Consume(context.Background(), uint(u), client, topic)
	return h
}
func (h *Hub) CreateGame(client *Client, topic string, token string, gameMessage *CreateGameMessage) *Hub {
	var tr models.TriviaService
	tr = models.S(gameMessage.Code)
	gameSession, created := tr.GetOrCreateGameSession(gameMessage.Code, gameMessage.TriviaID)
	var msg GameMessage
	msg.Type = JOINED
	if created {
		CreateTopic(1, topic, 1)
		msg.Type = CREATED
	}
	m, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	err = client.Send(m)
	if err != nil {
		return nil
	}
	h.Subscribe(client, gameSession.Code, token)
	return h
}

func (h *Hub) StartGame(client *Client, topic string, token string) *Hub {
	Countdown(5, topic)
	StartGameSession()
	return h
}

func (h *Hub) StartgameSession() {

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

func (h *Hub) CurrentUsers(topic string) {
	subscriptions := h.GetSubscriptions(topic, nil)
	var currentUsers []User
	for _, sub := range subscriptions {
		user, err := userServices.GetUserByID(int64(sub.UserId))
		if err != nil {
			log.Println(err)
		}
		u := User{Username: user.Username}
		currentUsers = append(currentUsers, u)
	}
	msg := CurrentUsersMessage{Users: currentUsers, Type: CURRENT_USERS}
	b, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	h.Publish(topic, b)
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
	case CREATE_OR_JOIN:
		cr := CreateGameMessage{}
		err := json.Unmarshal(m.Data, &cr)
		if err != nil {
			log.Println("Missing trivia ID.")
		}
		h.CreateGame(&client, m.Topic, m.Token, &cr)
		fmt.Printf("Starting a new game with code: %s", m.Topic)
	case START:
		fmt.Printf("Starting a created game with code: %s", m.Topic)
		h.StartGame(&client, m.Topic, m.Token)
	default:
		break
	}
	return h
}

func (client *Client) Send(message []byte) error {
	return client.Connection.WriteMessage(1, message)

}
