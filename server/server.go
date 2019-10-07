package main

import (
	"context"
	"log"
	"net"

	pb "github.com/glebaltshifter/grpc-test/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedGrpcTestServer
}

func (s *server) GetQuotient(ctx context.Context, in *pb.DivisionPair) (*pb.DivisionResult, error) {
	log.Printf("Received: %v, %v", in.Dividend, in.Divisor)
	return &pb.DivisionResult{Value: in.Dividend / in.GetDivisor()}, nil
}

func (s *server) GetRemainder(ctx context.Context, in *pb.DivisionPair) (*pb.DivisionResult, error) {
	log.Printf("Received: %v, %v", in.Dividend, in.Divisor)
	return &pb.DivisionResult{Value: in.Dividend % in.GetDivisor()}, nil
}

// func (s *server) StreamLambs(ctx context.Context, in *pb.StreamLambsRequest, stream pb.GrpcTest_StreamLambsServer) error {
// 	log.Printf("Received lambs quantity: %v", in.Quantity)
// 	log.Println("Starting to stream...")
// 	var i int32
// 	for i = 0; i < in.Quantity; i++ {
// 		var msg pb.LambsMessage
// 		msg.Content = "test"
// 		if err := stream.Send(&msg); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGrpcTestServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
