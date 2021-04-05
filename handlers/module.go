package handlers

import (
	"github.com/andrearampin/dependency-injection-fx/handlers/home"
	"go.uber.org/fx"
	"net/http"
)

func Register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/", h)
}

var Module = fx.Provide(
	home.NewHandler,
)
