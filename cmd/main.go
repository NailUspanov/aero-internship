package main

import (
	"aero-internship/gen/api"
	"aero-internship/internal/adapters/handlers/v1"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	srv := &v1.GRPCServer{}
	go func() {
		mux := runtime.NewServeMux()
		api.RegisterContentCheckServiceHandlerServer(context.Background(), mux, srv)

		log.Fatalln(http.ListenAndServe("localhost:8000", mux))
	}()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	api.RegisterContentCheckServiceServer(s, srv)

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
