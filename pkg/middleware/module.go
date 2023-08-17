package middleware

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Dependencies struct {
	fx.In

	Logger *logrus.Logger
}

type MWProvider struct {
	Logger *logrus.Logger
}

func NewMWProvider(params Dependencies) *MWProvider {
	return &MWProvider{
		Logger: params.Logger,
	}
}

var Module = fx.Provide(NewMWProvider)
