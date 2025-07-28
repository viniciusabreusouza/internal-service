package di

import (
	"context"

	http "example.com/internal-service/internal/infra/http"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	ctx        context.Context
	cancel     context.CancelFunc
	log        *zap.Logger
	httpServer http.Server
}

// GetLogger retorna o logger da aplicação
func (a Application) GetLogger() *zap.Logger {
	return a.log
}

func NewApplication(ctx context.Context, log *zap.Logger, server http.Server) Application {
	ctx, cancel := context.WithCancel(ctx)

	return Application{
		ctx:        ctx,
		cancel:     cancel,
		log:        log,
		httpServer: server,
	}
}

func (app Application) Run() error {
	errGroup, ctx := errgroup.WithContext(app.ctx)

	go func() {
		<-ctx.Done()
		app.ShutdownAndCleanup()
	}()

	errGroup.Go(func() error {
		return app.httpServer.Run(ctx)
	})

	app.log.Info("Application started")

	return errGroup.Wait()
}

func (app Application) ShutdownAndCleanup() {
	log := app.GetLogger()
	log.Info("Shutting down server...")

	app.cancel()

	if app.httpServer != nil {
		if err := app.httpServer.Stop(context.Background()); err != nil {
			log.Error("Failed to shutdown HTTP server", zap.Error(err))
		}
	}
}
