package kenobi_logger

import (
	"errors"
	"github.com/ereb-or-od/kenobi-logger/pkg/interfaces"
	"github.com/ereb-or-od/kenobi-logger/pkg/options"
	"github.com/ereb-or-od/kenobi-logger/pkg/providers"
	"github.com/ereb-or-od/kenobi-logger/pkg/utilities"
	"go.uber.org/zap"
)

type logger struct {
	logger *zap.Logger
}

func (l logger) Debug(msg string, parameters ...map[string]interface{}) {
	l.logger.Debug(msg, *utilities.ToZapFields(parameters...)...)
}

func (l logger) Warn(msg string, parameters ...map[string]interface{}) {
	l.logger.Warn(msg, *utilities.ToZapFields(parameters...)...)
}

func (l logger) Error(msg string, err error, parameters ...map[string]interface{}) {
	if err == nil {
		err = errors.New("an unrecognized error")
	}
	zapFields := *utilities.ToZapFields(parameters...)
	if zapFields == nil || len(zapFields) == 0 {
		zapFields = make([]zap.Field, 0)
		zapFields = append(zapFields, zap.Error(err))
	}
	l.logger.Error(msg, zapFields...)
}

func (l logger) Fatal(msg string, err error, parameters ...map[string]interface{}) {
	if err == nil {
		err = errors.New("an unrecognized error")
	}
	zapFields := *utilities.ToZapFields(parameters...)
	if zapFields == nil || len(zapFields) == 0 {
		zapFields = make([]zap.Field, 0)
		zapFields = append(zapFields, zap.Error(err))
	}
	l.logger.Fatal(msg, zapFields...)
}

func (l logger) Info(msg string, parameters ...map[string]interface{}) {
	l.logger.Info(msg, *utilities.ToZapFields(parameters...)...)
}

func New(loggerOptions ...*options.LoggerOptions) (interfaces.Logger, error) {
	zapLogger, err := providers.New(loggerOptions...)
	if err != nil {
		return nil, err
	}
	return &logger{
		logger: zapLogger,
	}, nil
}
