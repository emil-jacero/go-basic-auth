package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	v3 "github.com/emil-jacero/grpc-basic-auth/pkg/auth/v3"
	envoy_service_auth_v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 9001, "gRPC port")
	flag.Parse()

	username := os.Getenv("AUTH_USERNAME")
	password := os.Getenv("AUTH_PASSWORD")

	if username == "" || password == "" {
		log.Fatalf("environment variables AUTH_USERNAME and AUTH_PASSWORD must be set")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen to %d: %v", *port, err)
	}

	gs := grpc.NewServer()
	envoy_service_auth_v3.RegisterAuthorizationServer(gs, v3.New(username, password))

	log.Printf("starting gRPC server on: %d\n", *port)
	gs.Serve(lis)
}
