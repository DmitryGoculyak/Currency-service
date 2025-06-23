package service

import (
	"Currency-service/internal/db"
	"Currency-service/internal/db/models"
	"Currency-service/proto"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	var currency models.CurrencyDB
	err := db.DB.Get(&currency, "SELECT * FROM currencies WHERE currency_code = $1", req.Code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			return nil, status.Error(codes.NotFound, "Currency not found")
		}
		log.Fatal("Error:", err)
		return nil, err
	}
	return &proto.CurrencyResponse{
		Code: currency.CurrencyCode,
		Name: currency.CurrencyName,
	}, nil
}

func (s *CurrencyServer) GetListCurrencies(ctx context.Context, _ *proto.Empty) (*proto.ListCurrenciesResponse, error) {
	var rows []models.CurrencyDB
	err := db.DB.Select(&rows, "SELECT * FROM currencies")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Currency not found")
		}
		log.Fatal("Error:", err)
		return nil, err
	}
	var currencies []*proto.CurrencyResponse
	for _, r := range rows {
		currencies = append(currencies, &proto.CurrencyResponse{
			Code: r.CurrencyCode,
			Name: r.CurrencyName,
		})
	}
	return &proto.ListCurrenciesResponse{Currency: currencies}, nil
}

func (s *CurrencyServer) DeleteAllCurrency(ctx context.Context, _ *proto.Empty) (*proto.DeleteCurrencyResponse, error) {
	_, err := db.DB.Exec("DELETE FROM currencies")
	if err != nil {
		log.Println(err)
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
