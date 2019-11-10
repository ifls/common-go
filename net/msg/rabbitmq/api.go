package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var conn *amqp.Connection
var channel *amqp.Channel

var (
	uri          = "amqp://guest:guest@47.107.151.251:5672/"
	exchange     = "test-exchange"
	exchangeType = "direct"
	routingKey   = "test-key"
	queueKey     = "test-key"
	body         = "foobar"
	reliable     = true
	queue        = "test_queue"
)

func init() {
	connection, err := amqp.Dial(uri)
	if err != nil {
		log.Println(err)
	}
	conn = connection

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}

	channel = ch
}

func createExchange() error {
	if err := channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	return nil
}

func publish() error {
	//发布到exchange
	if err := channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}

	return nil
}

func createQueue() (*amqp.Queue, error) {
	queue, err := channel.QueueDeclare(
		queue, // name of the queue
		true,  // 存盘
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	return &queue, nil
}

func queueBind() error {
	if err := channel.QueueBind(
		queue,    // name of the queue
		queueKey, // bindingKey
		exchange, // sourceExchange
		false,    // noWait
		nil,      // arguments
	); err != nil {
		return fmt.Errorf("Queue Bind: %s", err)
	}

	return nil
}

func consumer() error {
	deliveries, err := channel.Consume(
		queue,    // name
		queueKey, // consumerTag,
		false,    // noAck
		false,    // exclusive
		false,    // noLocal
		false,    // noWait
		nil,      // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Consume: %s", err)
	}

	for {
		select {
		case d := <-deliveries:
			log.Printf("msg = %s\n", d.Body)
		}
	}

	return nil
}

func confirmOne(confirms <-chan amqp.Confirmation) {
	log.Printf("waiting for confirmation of one publishing")

	if confirmed := <-confirms; confirmed.Ack {
		log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		log.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}
