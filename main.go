package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/cosmotek/grpc-template/api"
	grpcapi "github.com/cosmotek/grpc-template/api/gen/proto/go/v1/api"

	"github.com/flowchartsman/swaggerui"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func runHTTP(server grpcapi.HelloWorldServer) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// register the handler
	err := grpcapi.RegisterHelloWorldHandlerServer(ctx, mux, server)
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Handle("/grpc/*", http.StripPrefix("/grpc", mux))
	r.Mount("/swagger", http.StripPrefix("/swagger", swaggerui.Handler(api.SwaggerSpecJSON)))

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", r)
}

func runGRPC(server grpcapi.HelloWorldServer) error {
	listener, err := net.Listen("tcp", ":5051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpcapi.RegisterHelloWorldServer(grpcServer, server)

	// register reflection service
	reflection.Register(grpcServer)

	// Start gRPC server
	return grpcServer.Serve(listener)
}

func main() {
	server := &HelloWorldServer{}

	log.Println("starting http service on :8081")
	log.Println("starting grpc service on :5051")
	log.Println("swaggerui available at http://localhost:8081/swagger/")

	// run both listeners (TODO improve error handling here...)
	go runGRPC(server)
	runHTTP(server)
}

type HelloWorldServer struct{}

func (h HelloWorldServer) Greet(ctx context.Context, input *grpcapi.Hello) (*grpcapi.Goodbye, error) {
	switch t := input.FormatRule.(type) {
	case *grpcapi.Hello_Language:
		return &grpcapi.Goodbye{Msg: wrapperspb.String(input.Msg.GetValue() + "-" + t.Language)}, nil
	case *grpcapi.Hello_PreventTranslations:
		return nil, nil
	default:
		return nil, status.Error(codes.InvalidArgument, "oneof 'format_rule' was empty or malformed")
	}
}

func (h HelloWorldServer) GreetOther(ctx context.Context, input *grpcapi.Hello) (*grpcapi.Goodbye, error) {
	switch t := input.FormatRule.(type) {
	case *grpcapi.Hello_Language:
		return &grpcapi.Goodbye{Msg: wrapperspb.String(input.Msg.GetValue() + "-" + t.Language)}, nil
	case *grpcapi.Hello_PreventTranslations:
		return nil, nil
	default:
		return nil, status.Error(codes.InvalidArgument, "oneof 'format_rule' was empty or malformed")
	}
}

func (h HelloWorldServer) GreetOther2(ctx context.Context, input *grpcapi.Hello) (*grpcapi.Goodbye, error) {
	switch t := input.FormatRule.(type) {
	case *grpcapi.Hello_Language:
		return &grpcapi.Goodbye{Msg: wrapperspb.String(input.Msg.GetValue() + "-" + t.Language)}, nil
	case *grpcapi.Hello_PreventTranslations:
		return nil, nil
	default:
		return nil, status.Error(codes.InvalidArgument, "oneof 'format_rule' was empty or malformed")
	}
}

func (h HelloWorldServer) GreetReverse(ctx context.Context, input *grpcapi.Goodbye) (*grpcapi.Goodbye, error) {
	return input, nil
}

func (h HelloWorldServer) OriginalGreet(ctx context.Context, input *grpcapi.Goodbye) (*grpcapi.Goodbye, error) {
	return input, nil
}
