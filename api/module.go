package handler

import (
	"github.com/FaridunM/Banks_credits_parser/pkg/config"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Dependencies struct {
	fx.In

	Logger  *logrus.Logger
	Credits *config.Credits
}

type Handler struct {
	logger  *logrus.Logger
	credits *config.Credits
}

func NewHandler(params Dependencies) *Handler {
	return &Handler{
		logger:  params.Logger,
		credits: params.Credits,
	}
}

var Module = fx.Provide(NewHandler)
