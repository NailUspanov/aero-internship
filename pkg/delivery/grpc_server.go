package delivery

import (
	pb "aero-internship/gen/api"
	v1 "aero-internship/internal/adapters/handlers/v1"
	"aero-internship/pkg/config"

	"google.golang.org/grpc"
)

// функция создает новый gRPC сервер
func NewGRPCServer(cfg *config.Config) (*grpc.Server, error) {
	grpc_server := grpc.NewServer()
	pb.RegisterContentCheckServiceServer(grpc_server, &v1.GRPCServer{})
	// pb.RegisterNewsServiceServer(grpc_server, &newsServiceServer{})
	// pb.RegisterTagServiceServer(grpc_server, &tagServiceServer{})
	return grpc_server, nil
}
