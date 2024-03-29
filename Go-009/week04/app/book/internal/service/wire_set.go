package service

import (
	"G20220607100065/Go-009/week04/app/book/internal/biz"
	"G20220607100065/Go-009/week04/app/book/internal/data/ent"
	"github.com/google/wire"
)

// ProvideSet for service package ...
var ProvideSet = wire.NewSet(NewHTTPBookRepo, NewGRPCBookRepo, NewGRPCBookService)

func NewHTTPBookRepo(client *ent.Client) biz.HTTPBookRepo {
	return &HTTPBookService{Client: client}
}

func NewGRPCBookRepo(client *ent.Client) biz.GRPCBookRepo {
	return &GRPCBookService{Client: client}
}

func NewGRPCBookService(client *ent.Client) *GRPCBookService {
	return &GRPCBookService{Client: client}
}

//https://github.com/qmdx00/Go-Camp/blob/master/G20220607100065/Go-009/week04/app/book/internal/service/wire_set.go
