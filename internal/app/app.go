package app

import (
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"aero-internship/pkg/client/postgres"
	"aero-internship/pkg/config"
	"aero-internship/pkg/delivery"
)

type App struct {
	cfg        *config.Config
	restServer *delivery.RESTServer
	grpcServer *grpc.Server
	db         *sqlx.DB
}

func NewApp(cfg *config.Config) (*App, error) {
	restServer, err := delivery.NewRESTServer(cfg)
	if err != nil {
		log.Fatalf("[REST] Can't create new REST server: %v", err)
	}
	grpcServer, err := delivery.NewGRPCServer(cfg)
	if err != nil {
		log.Fatalf("[gRPC] Can't create new gRPC server: %v", err)
	}
	//соединение с бд
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize db.sql %s", err.Error())
	}
	return &App{
		cfg:        cfg,
		restServer: restServer,
		grpcServer: grpcServer,
		db:         db,
	}, nil
}

// в отдельных горутинах запускаем gRPC сервер и REST шлюз
func (application *App) Run() {
	go func() {
		log.Printf(
			"[REST] Server listening at %s:%s",
			application.cfg.GetRESTHost(),
			application.cfg.GetRESTPort(),
		)
		application.restServer.Run(application.cfg)
	}()
	go func() {
		lis, err := net.Listen(
			"tcp",
			fmt.Sprintf(":%s", application.cfg.GetGRPCPort()),
		)
		if err != nil {
			log.Fatalf("[gRPC] Failed to listen: %v", err)
		}
		log.Printf("[gRPC] Server listening at %v", lis.Addr())
		if err := application.grpcServer.Serve(lis); err != nil {
			log.Fatalf("[gRPC] Failed to serve: %v", err)
		}
	}()
}