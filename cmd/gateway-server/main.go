package main

import (
	"context"
	"log"
	"net"

	"github.com/bmoffatt/protobuf-gateway-example/must"
	"github.com/bmoffatt/protobuf-gateway-example/pb"
	"google.golang.org/grpc"
)

const (
	myAddr     = ":9000"
	targetAddr = ":9001"
)

type gatewayServer struct {
	pb.UnimplementedGatewayServer
}

func (s *gatewayServer) DoTheThing(ctx context.Context, req *pb.GatewayRequest) (*pb.GatewayResponse, error) {
	log.Print(req)
	conn := must.Return(grpc.Dial(targetAddr, grpc.WithInsecure()))
	return pb.NewApplicationClient(conn).TheThing(ctx, &pb.ApplicationRequest{
		Id:    req.Id,
		Stuff: req.Stuff,
	})
}

func init() {
	log.SetPrefix("(gateway) ")
}

func main() {
	lis := must.Return(net.Listen("tcp", myAddr))
	log.Printf("gateway listening at %s", myAddr)
	s := grpc.NewServer()
	pb.RegisterGatewayServer(s, &gatewayServer{})
	must.Do(s.Serve(lis))
}
