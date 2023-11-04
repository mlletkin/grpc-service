//go:generate mockgen -source ./receiver.go -destination=./mocks/receiver.go -package=mock_logger
package logger

import (
	"github.com/IBM/sarama"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/infrastructure/kafka"
)

type HandleFunc func(msg *sarama.ConsumerMessage)

type Receiver interface {
	Subcribe(topic string) error
}
type KafkaReceiver struct {
	consumer *kafka.Consumer
	handlers map[string]HandleFunc
}

func NewKafkaReceiver(consumer *kafka.Consumer, handlers map[string]HandleFunc) Receiver {
	return &KafkaReceiver{
		consumer: consumer,
		handlers: handlers,
	}
}

func (r *KafkaReceiver) Subcribe(topic string) error {
	handler, ok := r.handlers[topic]
	if !ok {
		return errors.New("can not find handler")
	}
	partitions, err := r.consumer.SingleConsumer.Partitions(topic)
	if err != nil {
		return err
	}
	initialOffset := sarama.OffsetNewest

	for _, partition := range partitions {
		pc, err := r.consumer.SingleConsumer.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer, partition int32) {
			for message := range pc.Messages() {
				handler(message)
			}
		}(pc, partition)
	}
	return nil
}
