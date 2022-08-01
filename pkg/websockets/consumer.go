package websockets

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"strconv"
	"time"
)

func Consume(ctx context.Context, userId uint, client *Client, topic string) {
	l := log.New(os.Stdout, "kafka reader: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:         []string{brokerAddress},
		Topic:           topic,
		GroupID:         strconv.Itoa(int(userId)),
		Logger:          l,
		MinBytes:        1,
		MaxBytes:        10e6,
		MaxWait:         3 * time.Second,
		StartOffset:     kafka.FirstOffset,
		ReadLagInterval: -1,
	})
	defer func(r *kafka.Reader) {
		err := r.Close()
		if err != nil {

		}
	}(r)
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))

		err = client.Send(msg.Value)
		if err != nil {
			return
		}
	}
}
