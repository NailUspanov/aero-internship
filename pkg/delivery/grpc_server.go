package delivery

import (
	pb "aero-internship/gen/api"
	"aero-internship/internal/adapters/handlers"
	"aero-internship/internal/domain/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// функция создает новый gRPC сервер
func NewGRPCServer(handlers *handlers.Handler, services *usecase.Service) (*grpc.Server, error) {

	opts := []grpc.ServerOption{grpc.Creds(insecure.NewCredentials()), grpc.UnaryInterceptor(services.UnaryInterceptor)}
	grpc_server := grpc.NewServer(opts...)
	pb.RegisterContentCheckServiceServer(grpc_server, handlers.ContentCheckServiceServer)
	// pb.RegisterNewsServiceServer(grpc_server, &newsServiceServer{})
	// pb.RegisterTagServiceServer(grpc_server, &tagServiceServer{})
	return grpc_server, nil
}
