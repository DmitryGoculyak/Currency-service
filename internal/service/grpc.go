package service

import (
	"Currency-service/internal/db"
	"Currency-service/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type CurrencyServer struct {
	proto.UnimplementedCurrencyServiceServer
}

func (s *CurrencyServer) CreateCurrency(ctx context.Context, req *proto.CreateCurrencyRequest) (*proto.CurrencyResponse, error) {
	_, err := db.DB.Exec("INSERT INTO currencies(currency_code,currency_name) VALUES($1, $2)", req.Code, req.Name)
	if err != nil {
		log.Fatal(err)
	}
	return &proto.CurrencyResponse{Code: req.Code, Name: req.Name}, nil
}

func (s *CurrencyServer) GetCurrencies(ctx context.Context, req *proto.GetCurrenciesRequest) (*proto.CurrencyResponse, error) {
	var currency proto.CurrencyResponse
	err := db.DB.Get(&currency, "SELECT * FROM currencies WHERE currency_code = $1", req.Code)
	if err != nil {
		return nil, err
	}
	return &currency, nil
}

func (s *CurrencyServer) GetListCurrencies(ctx context.Context, _ *proto.Empty) (*proto.ListCurrenciesResponse, error) {
	var currencies []*proto.CurrencyResponse
	err := db.DB.Select(&currencies, "SELECT * FROM currencies")
	if err != nil {
		return nil, err
	}
	return &proto.ListCurrenciesResponse{Currency: currencies}, nil
}

func (s *CurrencyServer) DeleteAllCurrency(ctx context.Context, _ *proto.Empty) (*proto.DeleteCurrencyResponse, error) {
	_, err := db.DB.Exec("DELETE FROM currencies")
	if err != nil {
		return nil, err
	}
	return &proto.DeleteCurrencyResponse{Message: "All currency deleted"}, nil
}

func RunService() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterCurrencyServiceServer(grpcServer, &CurrencyServer{})

	log.Printf("[gRPC] Server started on port: %v", lis.Addr())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
