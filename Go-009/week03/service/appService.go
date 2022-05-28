package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// App Server生命周期托管对象
type App struct {
	ctx      context.Context
	cancel   func()
	logger   *log.Logger
	servers  []*http.Server
	errGroup *errgroup.Group
}

type AppOption func(app *App)

// NewApp 返回App对象
func NewApp(opts ...AppOption) *App {
	app := &App{
		ctx:     context.Background(),
		logger:  log.Default(),
		servers: make([]*http.Server, 0, 3),
	}

	for _, opt := range opts {
		opt(app)
	}

	app.ctx, app.cancel = context.WithCancel(app.ctx)
	app.errGroup, app.ctx = errgroup.WithContext(app.ctx)
	return app
}

// Run 开启服务
func (app *App) Run() error {
	app.startListenSystemSignal() // 开启系统信号监听
	app.startServers()            // 开启托管的服务
	return app.errGroup.Wait()    // 等待errGroup的结束信号
}

// Stop 关闭服务
func (app *App) Stop() error {
	app.logger.Println("begin to stop the app...")
	app.cancel()
	return app.errGroup.Wait()
}

// 开启托管的服务。
func (app *App) startServers() {
	// 为了处理闭包问题，包一层。
	wrapServerStart := func(srv *http.Server) func() error {
		go func() {
			<-app.ctx.Done()
			app.logger.Printf("begin to shutdown the server: %v\n", srv.Addr)
			if err := srv.Shutdown(app.ctx); err != nil {
				log.Printf("HTTP server %s Shutdown error: %v", srv.Addr, err)
			}
		}()

		return func() error {
			app.logger.Printf("start the server: %v\n", srv.Addr)
			return errors.WithMessagef(srv.ListenAndServe(), "server is exit: %s", srv.Addr)
		}
	}

	for _, srv := range app.servers {
		app.errGroup.Go(wrapServerStart(srv))
	}
}

// 开启系统信号监听
func (app *App) startListenSystemSignal() {
	app.errGroup.Go(func() error {
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-app.ctx.Done():
			signal.Stop(signalCh)
			close(signalCh)
			return app.ctx.Err()
		case s := <-signalCh:
			return errors.Errorf("receive linux's signal: %v", s)
		}
	})
}

func WithServer(server *http.Server) AppOption {
	return func(app *App) {
		app.servers = append(app.servers, server)
	}
}

func WithLog(logger *log.Logger) AppOption {
	return func(app *App) {
		app.logger = logger
	}
}
