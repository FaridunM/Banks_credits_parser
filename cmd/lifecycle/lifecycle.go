package lifecycle

import (
	"github.com/FaridunM/Banks_credits_parser/pkg/config"
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Register hooks for start or shutdown server
func RegisterHooks(lifecycle fx.Lifecycle, logger *logrus.Logger, config *config.Config, server *http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			url := fmt.Sprintf("%s:%s", config.Host, config.Port)
			logger.Info("Starting server... on " + url)
			go server.ListenAndServe()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	},
	)
}

// fx_lifecycle Module
var Module = fx.Invoke(RegisterHooks)
