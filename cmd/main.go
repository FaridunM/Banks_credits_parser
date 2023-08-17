package main

import (
	handler "github.com/FaridunM/Banks_credits_parser/api"
	"github.com/FaridunM/Banks_credits_parser/cmd/lifecycle"
	"github.com/FaridunM/Banks_credits_parser/cmd/router"
	"github.com/FaridunM/Banks_credits_parser/pkg/config"
	"github.com/FaridunM/Banks_credits_parser/pkg/logger"
	"github.com/FaridunM/Banks_credits_parser/pkg/middleware"
	"github.com/FaridunM/Banks_credits_parser/pkg/server"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		config.CreditModule,
		logger.Module,
		middleware.Module,
		handler.Module,
		router.Module,
		server.Module,
		lifecycle.Module,
		// cache.Module,		// cache.Module usually used redis
		// gateway.Module,
		// internal.Module,
		// fx.NopLogger,		// fx.NopLogger is used to disable native logging
	)
	app.Run()
}
