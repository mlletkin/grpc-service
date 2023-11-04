package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hwservice "gitlab.ozon.dev/kavkazov/homework-8/internal/hw_service"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/db"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository/postgresql"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/server"
	pb "gitlab.ozon.dev/kavkazov/homework-8/pkg/hw_service"
	"gitlab.ozon.dev/kavkazov/homework-8/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var brokers = []string{
	"127.0.0.1:9091",
	"127.0.0.1:9092",
	"127.0.0.1:9093",
}

func main() {
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
	defer done()

	var addr string
	flag.StringVar(&addr, "addr", ":50051", "address for homework_service server")

	errCh := make(chan error, 1)
	go func() {
		if err := run(ctx, addr); err != nil {
			errCh <- err
			close(errCh)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println()
		log.Println("shutting down...")

		done()
	case err := <-errCh:
		log.Println("exited with error:", err)

		done()
	}

}

func run(ctx context.Context, addr string) error {

	zapLogger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	logger.SetGlobal(
		zapLogger.With(zap.String("component", "homework_service")),
	)

	srv := grpc.NewServer()

	database, err := db.NewDB(ctx)
	if err != nil {
		return err
	}
	defer database.GetPool(ctx).Close()

	impl := server.NewServer(
		postgresql.NewPosts(database),
		postgresql.NewComments(database),
	)

	pb.RegisterHomeworkServiceServer(srv, hwservice.New(impl))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	logger.Infof(ctx, "homework service listening on %q", addr)

	return srv.Serve(lis)
}
