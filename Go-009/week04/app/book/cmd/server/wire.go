// +build wireinject

package main

import (
	"G20220607100065/Go-009/week04/app/book/internal/biz"
	"G20220607100065/Go-009/week04/app/book/internal/conf"
	"G20220607100065/Go-009/week04/app/book/internal/data"
	"G20220607100065/Go-009/week04/app/book/internal/data/ent"
	"G20220607100065/Go-009/week04/app/book/internal/server"
	"G20220607100065/Go-009/week04/app/book/internal/service"
	"github.com/google/wire"
	"golang.org/x/sync/errgroup"
)

type App struct {
	HttpServer *server.HttpServer
	GRPCServer *server.GRPCServer
	Client     *ent.Client
}

// newApp return App struct with server.HttpServer and server.GRPCServer
func newApp(http *server.HttpServer, grpc *server.GRPCServer, client *ent.Client) *App {
	return &App{HttpServer: http, GRPCServer: grpc, Client: client}
}

// initApp Inject wire ProvideSet
func initApp(group *errgroup.Group, option conf.Options) *App {
	panic(wire.Build(server.ProvideSet, data.ProvideSet, service.ProvideSet, biz.ProvideSet, newApp))
}
