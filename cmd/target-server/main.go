package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	"github.com/bmoffatt/protobuf-gateway-example/must"
	"github.com/bmoffatt/protobuf-gateway-example/pb/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	myAddr = ":9001"
)

type applicationServer struct {
	pb.UnimplementedApplicationServer
}

func (s *applicationServer) TheThing(ctx context.Context, req *pb.ApplicationRequest) (*pb.GatewayResponse, error) {
	log.Print(req.String())
	if rand.Int()%2 == 0 {
		return nil, must.Return(status.New(codes.DataLoss, "IDK").WithDetails(&pb.AnError{
			Yolo: "my BFF Jill",
		})).Err()
	}
	return &pb.GatewayResponse{
		Msg: "🤪",
	}, nil

}

func init() {
	log.SetPrefix("(app) ")
}

func main() {
	l := must.Return(net.Listen("tcp", myAddr))
	log.Printf("application listening at: %s", myAddr)
	s := grpc.NewServer()
	pb.RegisterApplicationServer(s, &applicationServer{})
	must.Do(s.Serve(l))
}
