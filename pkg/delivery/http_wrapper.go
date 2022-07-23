package delivery

import (
	"aero-internship/pkg/config"
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "aero-internship/gen/api"
)

type RESTServer struct {
	mux *runtime.ServeMux
}

func NewRESTServer(cfg *config.Config) (*RESTServer, error) {
	mux, err := NewWrapperMux(cfg)
	if err != nil {
		return nil, err
	}
	return &RESTServer{
		mux: mux,
	}, nil
}

func NewWrapperMux(cfg *config.Config) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()
	return mux, nil
}

func (restServer *RESTServer) Run(cfg *config.Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	pb.RegisterContentCheckServiceHandlerFromEndpoint(
		ctx,
		restServer.mux,
		fmt.Sprintf("%s:%s", cfg.GetGRPCHost(), cfg.GetGRPCPort()),
		opts,
	)
	if err := http.ListenAndServe(
		fmt.Sprintf(":%s", cfg.GetRESTPort()),
		restServer.mux,
	); err != nil {
		return err
	}
	return nil
}
