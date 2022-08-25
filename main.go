package main

import (
	"context"
	trippb "coolcar/server/proto/gen/go"
	trip "coolcar/server/tripservice"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", lis)
	}

	go startGRPCGateway()

	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	log.Fatal(s.Serve(lis))
}

func startGRPCGateway() {
	log.SetFlags(log.Lshortfile)

	c := context.Background()
	c, cancle := context.WithCancel(c)
	defer cancle()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions:   protojson.MarshalOptions{
				UseEnumNumbers: true,
				EmitUnpopulated: true,
				//UseProtoNames: true,
			},
			//UnmarshalOptions: protojson.UnmarshalOptions{
			//
			//},
		},
	))
	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, ":8081", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("cannot start grpc gateway: %v", err)
	}

	err = http.ListenAndServe(":8989", mux)
	if err != nil {
		log.Fatalf("cannot listen and server: %v", err)
	}

}