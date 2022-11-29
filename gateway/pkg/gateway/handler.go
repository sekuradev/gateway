package gateway

import (
	"context"
	"flag"
	"net/http"

	sekura "github.com/sekuradev/api/gateway/pkg/sekuraapi/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	grpcHost      = flag.String("h", "localhost:50053", "endpoint of the gRPC server")
	allowInsecure = flag.Bool("insecure", false, "use plaintext connection with the remote gRPC server")
)

func defaultVars() (context.Context, *runtime.ServeMux, []grpc.DialOption) {
	ctx := context.TODO()
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions:   protojson.MarshalOptions{UseProtoNames: true},
			UnmarshalOptions: protojson.UnmarshalOptions{DiscardUnknown: true},
		}),
	)
	opts := []grpc.DialOption{}
	if *allowInsecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return ctx, mux, opts
}

func AgentHandler() http.Handler {
	ctx, mux, opts := defaultVars()
	sekura.RegisterAgentServiceHandlerFromEndpoint(ctx, mux, *grpcHost, opts)
	return mux
}

func UIHandler() http.Handler {
	ctx, mux, opts := defaultVars()
	sekura.RegisterUIServiceHandlerFromEndpoint(ctx, mux, *grpcHost, opts)
	return mux
}

func AllHandler() http.Handler {
	ctx, mux, opts := defaultVars()
	sekura.RegisterAgentServiceHandlerFromEndpoint(ctx, mux, *grpcHost, opts)
	sekura.RegisterUIServiceHandlerFromEndpoint(ctx, mux, *grpcHost, opts)
	return mux
}
