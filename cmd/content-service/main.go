package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gqlgen-playground/internal/app/content"
	"gqlgen-playground/internal/config"
	contentv1 "gqlgen-playground/internal/pb/content/v1"
	content_service "gqlgen-playground/internal/pkg/service/content-service"
	content_service_storage "gqlgen-playground/internal/pkg/storage/content-service"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := ":50051"
	metricsAddr := ":9090"

	pgDSN := os.Getenv("CONTENT_PG_DSN")
	if pgDSN == "" {
		log.Fatal("CONTENT_PG_DSN env is required, e.g. postgres://user:pass@localhost:5432/content?sslmode=disable")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := config.ConnectToPostgres(ctx, pgDSN)
	if err != nil {
		log.Fatalf("failed to init postgres connection: %v", err)
	}
	defer conn.Close()

	storage, err := content_service_storage.NewStorage(ctx, conn)
	if err != nil {
		log.Fatalf("failed to init storage: %v", err)
	}

	contentService := content_service.NewService(storage)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", addr, err)
	}

	// 🔹 Регистрируем стандартные runtime-метрики Go и процесса
	if err := prometheus.Register(collectors.NewGoCollector()); err != nil {
		if _, ok := err.(prometheus.AlreadyRegisteredError); !ok {
			log.Fatalf("failed to register GoCollector: %v", err)
		}
	}

	if err := prometheus.Register(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})); err != nil {
		if _, ok := err.(prometheus.AlreadyRegisteredError); !ok {
			log.Fatalf("failed to register ProcessCollector: %v", err)
		}
	}

	// 🔹 Включаем измерение времени обработки RPC в виде гистограммы
	grpc_prometheus.EnableHandlingTimeHistogram()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		// если будут стримы — добавишь StreamInterceptor аналогично
	)

	// 🔹 Регистрируем gRPC-метрики (grpc_server_*), они пойдут в default registry
	grpc_prometheus.Register(grpcServer)

	contentv1.RegisterContentServiceServer(grpcServer, content.NewImplementation(contentService))

	// Опционально: удобно для grpcurl/grpcui
	reflection.Register(grpcServer)

	// 🔹 gRPC сервер
	go func() {
		log.Printf("content-service gRPC listening on %s\n", addr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// 🔹 HTTP /metrics
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())

		log.Printf("metrics HTTP endpoint listening on %s\n", metricsAddr)
		if err := http.ListenAndServe(metricsAddr, mux); err != nil {
			log.Fatalf("failed to serve metrics: %v", err)
		}
	}()

	// graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigCh
	log.Printf("received signal %s, shutting down gRPC server...", sig)
	grpcServer.GracefulStop()
	log.Println("server stopped")
}
