package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/infrastructure/kafka"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/logger"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"
)

func runTs(svr *mux.Router) httptest.Server {
	return *httptest.NewServer(svr)
}

func TestKafkaLoggerMiddleware(t *testing.T) {
	svr, ctrl := setUpServer(t)
	defer ctrl.Finish()

	ts := runTs(svr)

	rch := make(chan string)

	t.Run("consume messages", func(t *testing.T) {
		kafkaConsumer, err := kafka.NewConsumer(brokers)
		require.NoError(t, err)

		receiver := logger.NewKafkaReceiver(kafkaConsumer, map[string]logger.HandleFunc{
			"log": func(msg *sarama.ConsumerMessage) {
				lm := logger.LogMessage{}
				err = json.Unmarshal(msg.Value, &lm)
				if err != nil {
					log.Println("Consumer error:", err)
				}
				rch <- lm.Type
			},
		})

		receiver.Subcribe("log")

		consumer, err := kafka.NewConsumer(brokers)
		require.NoError(t, err)

		parts, err := consumer.SingleConsumer.Partitions("log")
		require.NoError(t, err)
		pc, err := consumer.SingleConsumer.ConsumePartition("log", parts[0], sarama.OffsetNewest)
		require.NoError(t, err)

		errch := make(chan error, 1)
		go func(pc sarama.PartitionConsumer, partition int32) {
			select {
			case <-time.After(5 * time.Second):
				errch <- errors.New("timeout")
				return
			case <-pc.Messages():
				return
			}
		}(pc, parts[0])

		response, err := ts.Client().Get(fmt.Sprintf("%s/post/?postID=1", ts.URL))

		require.NoError(t, err)
		require.NotEqual(t, http.StatusInternalServerError, response.StatusCode)

		select {
		case lastMsg := <-rch:
			assert.Equal(t, "GET", lastMsg)
		case err = <-errch:
			require.NoError(t, err)
		}
	})
}

func TestAddPost(t *testing.T) {
	ctx := context.Background()

	svr, ctrl := setUpServer(t)
	defer ctrl.Finish()

	ts := runTs(svr)
	t.Run("status ok", func(t *testing.T) {
		data, err := json.Marshal(&repository.Post{Text: "test", Heading: "test"})
		require.NoError(t, err)
		resp, err := ts.Client().
			Post(fmt.Sprintf("%s/post/", ts.URL), "application/json", bytes.NewReader(data))
		db.Truncate(ctx)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
