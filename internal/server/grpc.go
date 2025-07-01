package server

import (
	cfg "Currency-service/config"
	proto "Currency-service/pkg/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func RunService(cfg *cfg.GrpcServiceConfig, server proto.CurrencyServiceServer) {

	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterCurrencyServiceServer(grpcServer, server)

	log.Printf("[gRPC] Server started at time %v on port %v",
		time.Now().Format("[2006-01-02] [15:04]"), address)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
