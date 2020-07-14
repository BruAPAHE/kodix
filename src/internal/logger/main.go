package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func init() {
	var logger *zap.Logger

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("Cannot create logger")
	}

	Logger = logger.Sugar()
}
