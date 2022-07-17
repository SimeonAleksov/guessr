package websockets

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"strconv"
	"time"
)

const (
	topic         = "message-log"
	brokerAddress = "localhost:9092"
)

func Produce(ctx context.Context) {
	i := 0
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
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
		fmt.Println("writes:", i)
		i++
		time.Sleep(time.Second)
	}
}
