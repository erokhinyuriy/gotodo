package rabbitmq

import (
	"log"

	"github.com/wagslane/go-rabbitmq"
)

type rabbit struct {
	publisher *rabbitmq.Publisher
}

func NewPublisher() *rabbit {
	// error connection, need to fix
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("events"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	return &rabbit{publisher: publisher}
}

// w.i.p
func (r *rabbit) Publish(message string) error {

	err := r.publisher.Publish(
		[]byte(message),
		[]string{"my_routing_key"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("events"),
	)
	if err != nil {
		return err
	}

	return nil
}
