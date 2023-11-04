package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/infrastructure/kafka"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/db"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/logger"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/middleware"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository/postgresql"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/router"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/server"
)

const port = ":9000"

var brokers = []string{
	"127.0.0.1:9091",
	"127.0.0.1:9092",
	"127.0.0.1:9093",
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.NewDB(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer database.GetPool(ctx).Close()

	kafkaProducer, err := kafka.NewProducer(brokers)
	if err != nil {
		log.Fatalln(err)
	}

	kafkaConsumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		log.Fatalln(err)
	}

	sender := logger.NewKafkaSender(kafkaProducer, "log")

	receiver := logger.NewKafkaReceiver(kafkaConsumer, map[string]logger.HandleFunc{
		"log": func(msg *sarama.ConsumerMessage) {
			lm := logger.LogMessage{}
			err = json.Unmarshal(msg.Value, &lm)
			if err != nil {
				log.Println("Consumer error:", err)
			}
			fmt.Printf("[%s]   %v\n", string(msg.Key), lm)
		},
	})
	err = receiver.Subcribe("log")
	if err != nil {
		log.Fatalln("receiver error:", err)
	}

	implementation := server.NewServer(
		postgresql.NewPosts(database),
		postgresql.NewComments(database),
	)
	http.Handle("/", router.RootRouter(implementation, middleware.KafkaLogging(sender)))

	log.Println("server starts listening on port", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}
