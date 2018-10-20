package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a default "sugared" logger based on dev toggle
func NewLogger(dev bool) (sugar *zap.SugaredLogger, err error) {
	var logger *zap.Logger
	if dev {
		// Log:         DebugLevel
		// Encoder:     console
		// Errors:      stderr
		// Sampling:    no
		// Stacktraces: WarningLevel
		// Colors:      capitals
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = config.Build()
	} else {
		// Log:         InfoLevel
		// Encoder:     json
		// Errors:      stderr
		// Sampling:    yes
		// Stacktraces: ErrorLevel
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return
	}

	return logger.Sugar(), nil
}
