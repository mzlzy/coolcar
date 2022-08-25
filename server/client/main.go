package main

import (
	"context"
	trippb "coolcar/server/proto/gen/go"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect server: %v", err)
	}

	tsClient := trippb.NewTripServiceClient(conn)
	resp, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "456"})
	if err != nil {
		log.Fatalf("cannot call GetTrip: %v", err)
	}

	fmt.Println(resp)
}
