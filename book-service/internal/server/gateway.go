package server

import (
	"context"
	"net/http"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	desc "github.com/daniil4142/book-market/book-service/pkg/book-service"
)

func createGatewayServer(grpcAddr, gatewayAddr string, allowedOrigins []string) *http.Server {
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.

	conn, err := grpc.NewClient(
		grpcAddr,
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
			),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial server")
	}

	mux := runtime.NewServeMux()
	if err := desc.RegisterBookServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatal().Err(err).Msg("Failed registration handler")
	}

	gatewayServer := &http.Server{
		Addr:    gatewayAddr,
		Handler: cors(mux, allowedOrigins),
	}

	return gatewayServer
}

func cors(h http.Handler, allowedOrigins []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		providedOrigin := r.Header.Get("Origin")
		matches := false
		for _, allowedOrigin := range allowedOrigins {
			if providedOrigin == allowedOrigin {
				matches = true
				break
			}
		}

		if matches {
			w.Header().Set("Access-Control-Allow-Origin", providedOrigin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, grpc-metadata-log-level, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
