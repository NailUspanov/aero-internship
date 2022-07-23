package main

import (
	"aero-internship/gen/api"
	v1 "aero-internship/internal/adapters/handlers/v1"
	"aero-internship/internal/config"
	"aero-internship/pkg/client/postgres"
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// на этапе компиляции возникает слудующая ошибка: open .env: The system cannot find the file specified. Как я понял он не видит env. Нужна помощь, как указать путь к env?
	cfg := config.ConfigData()

	// инициализация env конфига
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//соединение с бд
	_, err := postgres.NewPostgresDB(postgres.Config{})

	if err != nil {
		logrus.Fatalf("failed to initialize db.sql %s", err.Error())
	}

	srv := &v1.GRPCServer{}
	go func() {
		mux := runtime.NewServeMux()
		api.RegisterContentCheckServiceHandlerServer(context.Background(), mux, srv)

		// подставляем переменные из конфига, только у нас localhost:8000, а в env он равен 5436, нужно ли в env добавлять еще порты??
		logrus.Fatalln(http.ListenAndServe(cfg.DB_HOST+cfg.DB_PORT, mux))
		// logrus.Fatalln(http.ListenAndServe("localhost:8000", mux))
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
