package handlers

import (
	"Currency-service/internal/service"
	proto "Currency-service/pkg/proto"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CurrencyHandler struct {
	proto.UnimplementedCurrencyServiceServer
	service service.CurrencyServiceServer
}

func CurrencyHandlerConstructor(
	service service.CurrencyServiceServer,
) *CurrencyHandler {
	return &CurrencyHandler{
		service: service,
	}
}

func (h *CurrencyHandler) CreateCurrency(ctx context.Context, req *proto.CreateCurrencyRequest) (*proto.CurrencyResponse, error) {
	code, err := h.service.Create(ctx, req.Code, req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.CurrencyResponse{
		Code: code.CurrencyCode,
		Name: code.CurrencyName,
	}, nil
}

func (h *CurrencyHandler) GetCurrencies(ctx context.Context, req *proto.GetCurrenciesRequest) (*proto.CurrencyResponse, error) {
	c, err := h.service.GetByCode(ctx, req.Code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Currency not found")
		}
		return nil, status.Error(codes.Internal, "Database error")
	}
	return &proto.CurrencyResponse{
		Code: c.CurrencyCode,
		Name: c.CurrencyName,
	}, nil
}

func (h *CurrencyHandler) GetListCurrencies(ctx context.Context, _ *proto.Empty) (*proto.ListCurrenciesResponse, error) {
	currencies, err := h.service.ListAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var list []*proto.CurrencyResponse
	for _, c := range currencies {
		list = append(list, &proto.CurrencyResponse{
			Code: c.CurrencyCode,
			Name: c.CurrencyName,
		})
	}
	return &proto.ListCurrenciesResponse{
		Currency: list,
	}, nil
}
