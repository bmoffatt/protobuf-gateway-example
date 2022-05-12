package main

import (
	"context"
	"log"

	"github.com/bmoffatt/protobuf-gateway-example/must"
	"github.com/bmoffatt/protobuf-gateway-example/pb/v2"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func init() {
	log.SetPrefix("(client) ")
}

func main() {
	log.Print("hello world!")
	conn := must.Return(grpc.Dial("localhost:9000", grpc.WithInsecure()))
	client := pb.NewGatewayClient(conn)
	response, err := client.DoTheThing(context.TODO(), &pb.GatewayRequest{
		Id: uuid.New().String(),
		Stuff: &pb.Stuff{
			Name:  "Bryan",
			Hobby: "demos",
			Age:   30,
		},
	})
	log.Print(response, err)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			log.Print("sad")
		}
		log.Print(st.Details()[0].(*pb.AnError))
	}
}
