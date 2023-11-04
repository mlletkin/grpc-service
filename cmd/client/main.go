package main

import (
	"context"
	"flag"
	"log"

	pb "gitlab.ozon.dev/kavkazov/homework-8/pkg/hw_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	flag.Parse()

	var addr string

	flag.StringVar(&addr, "addr", ":50051", "address to dial homework_service server")

	if err := run(ctx, addr); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, addr string) error {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := pb.NewHomeworkServiceClient(conn)

	post, err := client.GetPost(ctx, &pb.PostRequestWithId{Id: 1})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(post.GetEntity())

	return nil
}
