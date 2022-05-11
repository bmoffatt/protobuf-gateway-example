package main

import (
	"context"
	"log"

	"github.com/bmoffatt/protobuf-gateway-example/must"
	"github.com/bmoffatt/protobuf-gateway-example/pb/v2"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func init() {
	log.SetPrefix("(client) ")
}

func main() {
	log.Print("hello world!")
	conn := must.Return(grpc.Dial("localhost:9000", grpc.WithInsecure()))
	client := pb.NewGatewayClient(conn)
	log.Print(must.Return(client.DoTheThing(context.TODO(), &pb.GatewayRequest{
		Id: uuid.New().String(),
		Stuff: &pb.Stuff{
			Name:  "Bryan",
			Hobby: "demos",
			Age:   30,
		},
	})))
}
