package router

import (
	"net/http"

	handler "github.com/FaridunM/Banks_credits_parser/api"
	"github.com/FaridunM/Banks_credits_parser/pkg/middleware"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

func Router(h *handler.Handler, mw *middleware.MWProvider) (router *mux.Router) {
	router = mux.NewRouter().PathPrefix("/api").Subrouter()
	router.Use(mw.CorsMW(router))
	router.NotFoundHandler = http.HandlerFunc(h.Handle404)
	router.MethodNotAllowedHandler = http.HandlerFunc(h.Handle405)

	router.HandleFunc("/ping", h.Pong).Methods("GET", "OPTIONS")
	// router.HandleFunc("/getGredits", h.GetCredits).Methods("GET")
	router.HandleFunc("/getCreditsBy/{bank:[A-Za-z -]+}", h.GetCreditsBy).Methods("GET")
	router.HandleFunc("/getCredit/{bank:[A-Za-z -]+}/{credit_type}", h.GetCredit).Methods("GET")

	return
}

var Module = fx.Provide(Router)
