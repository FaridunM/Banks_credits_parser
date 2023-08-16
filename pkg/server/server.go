package server

import (
	"fmt"
	"net/http"

	"github.com/FaridunM/Banks_credits_parser/pkg/config"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

type Dependencies struct {
	fx.In

	Router *mux.Router
	Config *config.Config
}

// Builds the server from dependencies. Need lifecycle for success start the server
func Server(params Dependencies) *http.Server {
	url := fmt.Sprintf("%s:%s", params.Config.Host, params.Config.Port)

	return &http.Server{
		Addr:    url,
		Handler: params.Router,
	}
}

var Module = fx.Provide(Server)
