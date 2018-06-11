package lakeevent

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type EventChannel struct {
	consumer *kafka.Consumer
	producer *kafka.Producer
}

func NewEventChannel(kafka_ep, group_id, topic string) (*EventChannel, error) {
	// Kafka
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafka_ep,
		"group_id":          group_id,
	})

	if err != nil {
		return nil, err
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafka_ep,
	})
	if err != nil {
		return nil, err
	}

	return &EventChannel{consumer: consumer, producer: producer}, nil
}

func (ec *EventChannel) Send(message interface{}) error {
	return nil
}
