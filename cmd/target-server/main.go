package main

import (
	"context"
	"log"
	"net"

	"github.com/bmoffatt/protobuf-gateway-example/must"
	"github.com/bmoffatt/protobuf-gateway-example/pb/v2"
	"google.golang.org/grpc"
)

const (
	myAddr = ":9001"
)

type applicationServer struct {
	pb.UnimplementedApplicationServer
}

func (s *applicationServer) TheThing(ctx context.Context, req *pb.ApplicationRequest) (*pb.GatewayResponse, error) {
	log.Print(req.String())
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
