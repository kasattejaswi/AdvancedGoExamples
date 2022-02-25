package main

import (
	"log"
	"net"

	"github.com/kasatejaswi/grpcserverex/calender"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calender.RegisterCalenderServiceServer(s, &calender.CalenderServer{})
	calender.RegisterGreeterServiceServer(s, &calender.GreeterServer{})
	log.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
