package websockets

import (
	"context"
	"github.com/segmentio/kafka-go"
	"strconv"
	"time"
)

const (
	brokerAddress = "localhost:9092"
)

func Produce(ctx context.Context, code string) {
	i := 0
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        code,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: 1,
	}

	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
		i++
		time.Sleep(5 * time.Second)
	}
}
