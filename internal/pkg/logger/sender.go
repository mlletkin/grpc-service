package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/infrastructure/kafka"
)

type LogMessage struct {
	Type      string
	Path      string
	TimeStamp time.Time
}

type KafkaSender struct {
	producer *kafka.Producer
	topic    string
}

type Sender interface {
	buildMessage(msg LogMessage) (*sarama.ProducerMessage, error)
	SendAsyncMessage(msg LogMessage) error
	SendMessage(msg LogMessage) error
}

func NewKafkaSender(producer *kafka.Producer, topic string) *KafkaSender {
	return &KafkaSender{
		producer: producer,
		topic:    topic,
	}
}

func (ks *KafkaSender) buildMessage(msg LogMessage) (*sarama.ProducerMessage, error) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Println("Build message marshal error:", err)
		return nil, err
	}
	return &sarama.ProducerMessage{
		Topic:     ks.topic,
		Value:     sarama.ByteEncoder(jsonMsg),
		Partition: -1,
		Key:       sarama.StringEncoder(fmt.Sprint(msg.TimeStamp)),
	}, err
}

func (ks *KafkaSender) SendAsyncMessage(msg LogMessage) error {
	kafkaMsg, err := ks.buildMessage(msg)
	if err != nil {
		log.Println("Async send message marshal error:", err)
		return err
	}

	ks.producer.SendAsyncMessage(kafkaMsg)
	return nil
}

func (ks *KafkaSender) SendMessage(msg LogMessage) error {
	kafkaMsg, err := ks.buildMessage(msg)
	if err != nil {
		log.Println("Send message marshal error:", err)
		return err
	}
	_, _, err = ks.producer.SendSyncMessage(kafkaMsg)
	if err != nil {
		log.Println("Send message connector error:", err)
		return err
	}
	return nil
}
