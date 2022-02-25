package calender

import (
	"context"
	"log"
	"time"
)

type CalenderServer struct {
	UnimplementedCalenderServiceServer
}

type GreeterServer struct {
	UnimplementedGreeterServiceServer
}

func (s *CalenderServer) GetDate(ctx context.Context, in *DateRequest) (*DateResponse, error) {
	log.Printf("Got request for date")
	return &DateResponse{
		Date: time.Now().String(),
	}, nil
}

func (g *GreeterServer) GetGreeting(ctx context.Context, in *GreetingRequest) (*GreetingResponse, error) {
	return &GreetingResponse{
		Greeting: "Hello " + in.GetName(),
	}, nil
}
