package main

import (
	"aero-internship/gen/api"
	"aero-internship/internal/adapters/handlers/v1"
	"aero-internship/pkg/client/postgres"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
)

func main() {

	// инициализация env конфига
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//соединение с бд
	_, err := postgres.NewPostgresDB(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db.sql %s", err.Error())
	}

	srv := &v1.GRPCServer{}
	go func() {
		mux := runtime.NewServeMux()
		api.RegisterContentCheckServiceHandlerServer(context.Background(), mux, srv)
		logrus.Fatalln(http.ListenAndServe("localhost:8000", mux))
	}()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatal(err)
	}

	s := grpc.NewServer()

	api.RegisterContentCheckServiceServer(s, srv)

	if err := s.Serve(l); err != nil {
		logrus.Fatal(err)
	}
}
