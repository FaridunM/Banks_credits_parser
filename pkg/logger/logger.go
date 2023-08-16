package logger

import (
	"github.com/FaridunM/Banks_credits_parser/pkg/config"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Log *logrus.Logger

func Init(conf *config.Config) *logrus.Logger {
	if Log != nil {
		return Log
	}
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "../logs/info.log",
		logrus.ErrorLevel: "../logs/errors.log",
		logrus.WarnLevel:  "../logs/warn.log",
		logrus.PanicLevel: "../logs/panic.log",
	}

	Log = logrus.New()
	Log.SetReportCaller(true)
	Log.Hooks.Add(lfshook.NewHook(pathMap, &logrus.JSONFormatter{}))

	return Log
}

var Module = fx.Provide(Init)
