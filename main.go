package main

import (
	"context"
	"fmt"

	"github.com/thisPeyman/go-urlshortner/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewShortenerServiceClient(conn)

	res, err := client.ShortenUrl(context.Background(), &api.ShortenURLRequest{LongUrl: "https://google.com"})

	fmt.Println(res)
}
