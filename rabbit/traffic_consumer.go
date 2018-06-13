package rabbit

import (
	"github.com/streadway/amqp"
	"log"
	"fmt"
)

const queueName  = "traffic.write"

type Config struct {
	ConnectionString string
}

func Init(cfg Config) (<-chan  amqp.Delivery, error){
	conn, err := amqp.Dial(cfg.ConnectionString)
	fmt.Printf(cfg.ConnectionString)
	if err != nil {
		log.Fatal( "Failed to connect to RabbitMQ")
		return nil, err
	}


	ch, err := conn.Channel()
	if err != nil {
		log.Fatal( "Failed to open Channel")
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal( "Failed to declaire Queue")
		return nil, err
	}

	err = ch.Qos(
		1,
		0,
		false,
	)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal( "Failed to register Consumer")
		return nil, err
	}

	return msgs, nil
}
