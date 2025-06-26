package service

import (
	cfg "Currency-service/config"
	"Currency-service/internal/db/models"
	proto2 "Currency-service/pkg/proto"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

type CurrencyServer struct {
	proto2.UnimplementedCurrencyServiceServer
	db *sqlx.DB
}

func CurrencyServerContainer(
	db *sqlx.DB,
) *CurrencyServer {
	return &CurrencyServer{db: db}
}

func (s *CurrencyServer) CreateCurrency(ctx context.Context, req *proto2.CreateCurrencyRequest) (*proto2.CurrencyResponse, error) {
	_, err := s.db.Exec("INSERT INTO currencies(currency_code,currency_name) VALUES($1, $2)", req.Code, req.Name)
	if err != nil {
		return nil, err
	}
	return &proto2.CurrencyResponse{Code: req.Code, Name: req.Name}, nil
}

func (s *CurrencyServer) GetCurrencies(ctx context.Context, req *proto2.GetCurrenciesRequest) (*proto2.CurrencyResponse, error) {
	var currency models.CurrencyDB
	err := s.db.Get(&currency, "SELECT * FROM currencies WHERE currency_code = $1", req.Code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			return nil, status.Error(codes.NotFound, "Currency not found")
		}
		log.Fatal("Error:", err)
		return nil, err
	}
	return &proto2.CurrencyResponse{
		Code: currency.CurrencyCode,
		Name: currency.CurrencyName,
	}, nil
}

func (s *CurrencyServer) GetListCurrencies(ctx context.Context, _ *proto2.Empty) (*proto2.ListCurrenciesResponse, error) {
	var rows []models.CurrencyDB
	err := s.db.Select(&rows, "SELECT * FROM currencies")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Currency not found")
		}
		log.Fatal("Error:", err)
		return nil, err
	}
	var currencies []*proto2.CurrencyResponse
	for _, r := range rows {
		currencies = append(currencies, &proto2.CurrencyResponse{
			Code: r.CurrencyCode,
			Name: r.CurrencyName,
		})
	}
	return &proto2.ListCurrenciesResponse{Currency: currencies}, nil
}

func RunService(cfg *cfg.GrpcServiceConfig, server *CurrencyServer) {

	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto2.RegisterCurrencyServiceServer(grpcServer, server)

	log.Printf("[gRPC] Server started at time %v on port %v",
		time.Now().Format("[2006-01-02] [15:04]"), address)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
