package server

import (
	"G20220607100065/Go-009/week04/app/book/internal/conf"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type HttpServer struct {
	gp     *errgroup.Group
	config conf.HTTPOptions
	Server *http.Server
}

// NewHttpServer provide GRPCServer with conf.Options and gin.Engine ...
func NewHttpServer(group *errgroup.Group, options conf.Options, router *gin.Engine) *HttpServer {
	return &HttpServer{gp: group, config: options.Server.HTTP, Server: &http.Server{
		Handler:      router,
		Addr:         options.Server.HTTP.Addr,
		ReadTimeout:  options.Server.HTTP.Timeout,
		WriteTimeout: options.Server.HTTP.Timeout,
	}}
}

func (h *HttpServer) Serve(ctx context.Context) {
	h.Server.Addr = h.config.Addr
	h.gp.Go(func() error {
		fmt.Println("[BOOK API] http listen on", h.config.Addr)
		return h.Server.ListenAndServe()
	})
	h.gp.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("[BOOK API] http server exit ...")
			return h.Server.Shutdown(context.TODO())
		}
	})
}

// Stop stop the HttpServer ...
func (h *HttpServer) Stop() {
	h.gp.Go(func() error {
		fmt.Println("[BOOK API] http server stop ...")
		return h.Server.Shutdown(context.TODO())
	})
}
