package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	hwservice "gitlab.ozon.dev/kavkazov/homework-8/internal/hw_service"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/db"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository/postgresql"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/server"
	hw_service "gitlab.ozon.dev/kavkazov/homework-8/pkg/hw_service"
	"gitlab.ozon.dev/kavkazov/homework-8/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
	defer done()

	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger.SetGlobal(
		zapLogger.With(zap.String("component", "homework_service")),
	)

	var addr string
	flag.StringVar(&addr, "addr", ":50051", "address for homework_service server")

	errCh := make(chan error, 1)
	go func() {
		if err := run(ctx, addr); err != nil {
			errCh <- err

		}
	}()

	go func() {
		if err := runHttp(ctx, addr); err != nil {
			errCh <- err

		}
	}()

	select {
	case <-ctx.Done():
		logger.Infof(ctx, "shutting down...")
		done()
	case err := <-errCh:
		logger.Infof(ctx, "exited with error:", err)
		done()
	}

}

func runHttp(ctx context.Context, addr string) error {

	m := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := hw_service.RegisterHomeworkServiceHandlerFromEndpoint(ctx, m, fmt.Sprintf("localhost%s", addr), opts)

	if err != nil {
		logger.Errorf(ctx, "%v", err)

	}

	logger.Infof(ctx, "grpc server listening on :\"9000\"")

	return http.ListenAndServe(":9000", m)

}

func run(ctx context.Context, addr string) error {

	jaegerCfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	tracer, closer, err := jaegerCfg.New(
		"homework_service",
	)
	if err != nil {
		return fmt.Errorf("cannot create tracer: %v", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

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

	hw_service.RegisterHomeworkServiceServer(srv, hwservice.New(impl))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	logger.Infof(ctx, "grpc server listening on %q", addr)

	return srv.Serve(lis)
}
