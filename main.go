package main

import (
	"context"
	"github.com/andrearampin/dependency-injection-fx/clients"
	"github.com/andrearampin/dependency-injection-fx/handlers"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

func NewMux(lc fx.Lifecycle, logger *zap.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Sync()
			return server.Shutdown(ctx)
		},
	})

	return mux
}

var _commonComponents = fx.Options(
	handlers.Module,
)

func main() {
	app := fx.New(
		// Provide all the constructors we need, which teaches Fx how we'd like to
		// construct the *zap.Logger, http.Handler, and *http.ServeMux types.
		// Remember that constructors are called lazily, so this block doesn't do
		// much on its own.
		fx.Provide(
			NewLogger,
			NewMux,
			clients.NewDB,
		),
		_commonComponents,
		// Since constructors are called lazily, we need some invocations to
		// kick-start our application. In this case, we'll use Register. Since it
		// depends on an http.Handler and *http.ServeMux, calling it requires Fx
		// to build those types using the constructors above. Since we call
		// NewMux, we also register Lifecycle hooks to start and stop an HTTP
		// server.
		fx.Invoke(handlers.Register),
	)

	app.Run()
}
