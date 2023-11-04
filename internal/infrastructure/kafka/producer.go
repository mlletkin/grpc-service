package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

type Producer struct {
	brokers []string
	sync    sarama.SyncProducer
	async   sarama.AsyncProducer
}

func newSync(brokers []string) (sarama.SyncProducer, error) {
	syncProducerConfig := sarama.NewConfig()

	syncProducerConfig.Producer.Partitioner = sarama.NewRandomPartitioner

	syncProducerConfig.Producer.RequiredAcks = sarama.WaitForAll

	syncProducerConfig.Producer.Idempotent = true
	syncProducerConfig.Net.MaxOpenRequests = 1

	syncProducerConfig.Producer.CompressionLevel = sarama.CompressionLevelDefault

	syncProducerConfig.Producer.Return.Successes = true
	syncProducerConfig.Producer.Return.Errors = true

	syncProducerConfig.Producer.Compression = sarama.CompressionGZIP

	syncProducer, err := sarama.NewSyncProducer(brokers, syncProducerConfig)
	if err != nil {
		return nil, errors.Wrap(err, "error with sync kafka-producer")
	}

	return syncProducer, nil
}

func newAsync(brokers []string) (sarama.AsyncProducer, error) {
	asyncProducerConfig := sarama.NewConfig()

	asyncProducerConfig.Producer.Partitioner = sarama.NewHashPartitioner
	asyncProducerConfig.Producer.RequiredAcks = sarama.WaitForAll

	asyncProducerConfig.Producer.Return.Successes = true
	asyncProducerConfig.Producer.Return.Errors = true

	asyncProducer, err := sarama.NewAsyncProducer(brokers, asyncProducerConfig)
	if err != nil {
		return nil, errors.Wrap(err, "error with async kafka-producer")
	}

	go func() {
		for e := range asyncProducer.Errors() {
			fmt.Println(e.Error())
		}
	}()

	go func() {
		for m := range asyncProducer.Successes() {
			fmt.Println("Async success with key", m.Key)
		}
	}()

	return asyncProducer, nil
}

func NewProducer(brokers []string) (*Producer, error) {
	syncProducer, err := newSync(brokers)
	if err != nil {
		return nil, errors.Wrap(err, "error with sync kafka-producer")
	}

	asyncProducer, err := newAsync(brokers)
	if err != nil {
		return nil, errors.Wrap(err, "error with async kafka-producer")
	}

	producer := &Producer{
		brokers: brokers,
		sync:    syncProducer,
		async:   asyncProducer,
	}

	return producer, nil
}

func (k *Producer) Close() error {
	err := k.sync.Close()
	if err != nil {
		return errors.Wrap(err, "kafka.Connector.Close")
	}

	err = k.sync.Close()
	if err != nil {
		return errors.Wrap(err, "kafka.Connector.Close")
	}

	return nil
}

func (k *Producer) SendSyncMessage(
	message *sarama.ProducerMessage,
) (partition int32, offset int64, err error) {
	return k.sync.SendMessage(message)
}

func (k *Producer) SendAsyncMessage(message *sarama.ProducerMessage) {
	k.async.Input() <- message
}
