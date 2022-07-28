package delivery

import (
	pb "aero-internship/gen/api"
	"aero-internship/internal/adapters/handlers"
	"aero-internship/pkg/config"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type RESTServer struct {
	mux *runtime.ServeMux
}

func NewRESTServer(cfg *config.Config, tm *handlers.Handler) (*RESTServer, error) {
	mux, err := NewWrapperMux(cfg, *tm)
	if err != nil {
		return nil, err
	}
	return &RESTServer{
		mux: mux,
	}, nil
}

func NewWrapperMux(cfg *config.Config, s handlers.Handler) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(customMatcher),
	)

	mux.HandlePath("POST", "/signup", s.SignUp)
	mux.HandlePath("GET", "/signin", s.SignIn)
	mux.HandlePath("POST", "/test", s.Create)

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

func customMatcher(key string) (string, bool) {
	switch key {
	case "Authorization":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
