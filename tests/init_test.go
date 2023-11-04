package tests

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/infrastructure/kafka"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/logger"
	mock_logger "gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/logger/mocks"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/middleware"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository/postgresql"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/router"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/server"
	"gitlab.ozon.dev/kavkazov/homework-8/tests/postgres"
)

var db *postgres.TDB

var brokers = []string{
	"127.0.0.1:9091",
	"127.0.0.1:9092",
	"127.0.0.1:9093",
}

func setUpServer(t *testing.T) (*mux.Router, *gomock.Controller) {
	db = postgres.NewFromEnv()
	postRepo := postgresql.NewPosts(db.DB)

	ctrl := gomock.NewController(t)

	commentsRepo := postgresql.NewComments(db.DB)
	kafkaProducer, err := kafka.NewProducer(brokers)
	if err != nil {
		fmt.Println(err)
	}

	sender := logger.NewKafkaSender(kafkaProducer, "log")

	receiver := mock_logger.NewMockReceiver(ctrl)
	receiver.EXPECT().Subcribe("log").Return(nil)

	err = receiver.Subcribe("log")
	if err != nil {
		fmt.Println("receiver error:", err)
	}
	impl := server.NewServer(postRepo, commentsRepo)
	return router.RootRouter(impl, middleware.KafkaLogging(sender)), ctrl
}
