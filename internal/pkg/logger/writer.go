package logger

import "github.com/IBM/sarama"

type KafkaWriter struct {
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaWriter(producer sarama.SyncProducer, topic string) *KafkaWriter {

	return &KafkaWriter{producer: producer, topic: topic}
}

func (w *KafkaWriter) Write(p []byte) (n int, err error) {
	msg := &sarama.ProducerMessage{
		Topic: w.topic,
		Value: sarama.ByteEncoder(p),
	}
	_, _, err = w.producer.SendMessage(msg)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
