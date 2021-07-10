package main

import (
	"context"
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"pancake/maker/gen/api"
	"pancake/maker/handler"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	grpcZap.ReplaceGrpcLoggerV2(zapLogger)

	server := grpc.NewServer(grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(grpcZap.UnaryServerInterceptor(zapLogger), grpcAuth.UnaryServerInterceptor(auth))))
	api.RegisterPancakeBakerServiceServer(server, handler.NewBakerHandler())
	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}

func auth(ctx context.Context) (context.Context, error) {
	token, err := grpcAuth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	if token != "hi/mi/tsu" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid bearer token")
	}

	return context.WithValue(ctx, "UserName", "God"), nil
}
