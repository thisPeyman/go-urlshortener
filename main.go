package main

import (
	"context"
	"fmt"

	"github.com/thisPeyman/go-urlshortner/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewIDGeneratorServiceClient(conn)

	res, err := client.GenerateID(context.Background(), &api.GenerateIDRequest{})

	fmt.Println(res)
}
