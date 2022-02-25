package main

import (
	"context"
	"log"
	"time"

	"github.com/kasattejaswi/grpcclientex/calender"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	calenderClient := calender.NewCalenderServiceClient(conn)
	greeterClient := calender.NewGreeterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cr, err := calenderClient.GetDate(ctx, &calender.DateRequest{Date: ""})
	if err != nil {
		log.Fatal(err)
	}
	gr, err := greeterClient.GetGreeting(ctx, &calender.GreetingRequest{Name: "Tejaswi"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Date call response: ",cr.Date)
	log.Println("Greeting call response: ",gr.Greeting)
}
