package websockets

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
)

const (
	brokerAddress = "localhost:9092"
)

func Produce(ctx context.Context, code string, message ProducerMessage) {
	i := 0
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        code,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: 1,
	}

	payload, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	err = w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(i)),
		Value: payload,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}
