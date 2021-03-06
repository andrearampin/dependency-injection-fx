package home

import (
	"encoding/json"
	"github.com/andrearampin/dependency-injection-fx/clients"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type Welcome struct {
	Message string `json:"message"`
}

type Params struct {
	fx.In

	DB     clients.DB
	Logger *zap.Logger
}

func NewHandler(p Params) (http.Handler, error) {
	p.Logger.Info("Executing NewHandler.")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.Logger.Info("request new name")

		msg := Welcome{p.DB.Get()}

		js, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}), nil
}
