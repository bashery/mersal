package hub

import (
	"fmt"
	"hub"
	"strconv"
	"testing"

	"github.com/gorilla/websocket"
)

// test Topics.Subscribe() method
func TestSubscribe(t *testing.T) {

}

func BenchmarkSubscribe(b *testing.B) {}

func Testfunc() {
	var conn *websocket.Conn
	// test Subscribe
	for i := 0; i < 10; i++ {
		topic := "topic-" + strconv.Itoa(i)
		for i := 0; i < 100; i++ {
			hub.Subscribe(topic, conn)

		}

	}

	for i := 0; i < 10; i++ {
		topic := "topic-" + strconv.Itoa(i)
		fmt.Println("start Publishe to all Subscriber in", topic)

		hub.Publish(topic)
	}

}
