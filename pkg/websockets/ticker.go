package websockets

import (
	"context"
	"strconv"
	"time"
)

func Countdown(n int, topic string) {
	for n >= 0 {
		msg := ProducerMessage{
			Action: TICKER,
			Data:   strconv.Itoa(n),
		}
		Produce(context.Background(), topic, msg)
		time.Sleep(time.Second)
		n--
	}
}
