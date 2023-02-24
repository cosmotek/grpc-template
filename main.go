package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cosmotek/grpc-template/api"
	grpcapi "github.com/cosmotek/grpc-template/api/gen/proto/go/v1/api"

	"github.com/flowchartsman/swaggerui"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	server := &HelloWorldServer{}

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

func main() {
	log.Println("starting service on :8081")
	log.Println("swaggerui available at http://localhost:8081/swagger/")

	if err := run(); err != nil {
		log.Fatal(err)
	}
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
