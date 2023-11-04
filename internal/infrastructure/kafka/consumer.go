package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

type Consumer struct {
	brokers        []string
	SingleConsumer sarama.Consumer
}

func NewConsumer(brokers []string) (*Consumer, error) {
	cfg := sarama.NewConfig()

	cfg.Consumer.Return.Errors = true
	cfg.Consumer.Offsets.AutoCommit.Enable = true
	cfg.Consumer.Offsets.AutoCommit.Interval = 5 * time.Second

	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer(brokers, cfg)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		brokers:        brokers,
		SingleConsumer: consumer,
	}, err
}
