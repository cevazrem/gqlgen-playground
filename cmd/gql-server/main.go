package main

import (
	"gqlgen-playground/internal/app/graph"
	"gqlgen-playground/internal/app/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultPort = "8080"
const defaultContentServiceAddr = "localhost:50051"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	contentServiceAddr := os.Getenv("CONTENT_SERVICE_ADDR")
	if contentServiceAddr == "" {
		contentServiceAddr = defaultContentServiceAddr
	}

	// gRPC коннект к content-service
	conn, err := grpc.Dial(
		contentServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial content-service at %s: %v", contentServiceAddr, err)
	}
	defer conn.Close()

	resolver := graph.NewResolver(conn)

	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers:  resolver,
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	})

	graphQLServer := handler.New(schema)
	graphQLServer.AddTransport(transport.GET{})
	graphQLServer.AddTransport(transport.POST{})
	graphQLServer.AddTransport(transport.Options{})
	graphQLServer.AddTransport(transport.MultipartForm{})
	graphQLServer.Use(extension.Introspection{})

	// Playground и /query endpoint
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", graphQLServer)
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("🚀 GraphQL server running at http://localhost:%s/ (content-service: %s)", port, contentServiceAddr)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
