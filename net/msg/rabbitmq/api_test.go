package rabbitmq

import (
	"log"
	"testing"
)

func TestQueue(t *testing.T) {
	que, err := createQueue()
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("queue =%#v\n", que)
}

func TestExchange(t *testing.T) {
	err := createExchange()
	if err != nil {
		log.Println(err)
	}
}

func TestConsume(t *testing.T) {
	err := consumer()
	if err != nil {
		log.Println(err)
	}
}

func TestPublish(t *testing.T) {
	_ = createExchange()
	for i := 0; i < 99; i++ {
		err := publish()
		if err != nil {
			log.Println(err)
		}
	}
}

func TestR(t *testing.T) {
	_ = queueBind()
	confirmOne(nil)
}
