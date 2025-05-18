package rabbitmq

import (
	"log"

	"github.com/wagslane/go-rabbitmq"
)

type rabbit struct {
	publisher *rabbitmq.Publisher
}

func NewPublisher() *rabbit {
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@rabbitmq:5672/",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	//defer conn.Close()

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("logs"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &rabbit{publisher: publisher}
}

func (r *rabbit) Publish(message string) error {

	err := r.publisher.Publish(
		[]byte(message),
		[]string{"black"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("logs"),
	)

	if err != nil {
		return err
	}

	return nil
}
