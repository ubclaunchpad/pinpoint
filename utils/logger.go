package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a default "sugared" logger based on dev toggle. A non-blank
// logpath will use the json encoder and write to the provided filepath.
func NewLogger(dev bool, logpath string) (sugar *zap.SugaredLogger, err error) {
	var logger *zap.Logger
	var config zap.Config
	if dev {
		// Log:         DebugLevel
		// Encoder:     console
		// Errors:      stderr
		// Sampling:    no
		// Stacktraces: WarningLevel
		// Colors:      capitals
		config = zap.NewDevelopmentConfig()
		if logpath == "" {
			config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}
	} else {
		// Log:         InfoLevel
		// Encoder:     json
		// Errors:      stderr
		// Sampling:    yes
		// Stacktraces: ErrorLevel
		config = zap.NewProductionConfig()
	}

	// set log output configuration if provided
	if logpath != "" {
		if err = os.MkdirAll(filepath.Dir(logpath), os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create directories for logpath '%s': %s",
				logpath, err.Error())
		}
		config.OutputPaths = append(config.OutputPaths, logpath)
		config.Encoding = "json"
	}

	// instantiate logger
	if logger, err = config.Build(); err != nil {
		return nil, fmt.Errorf("new logger: %s", err.Error())
	}

	return logger.Sugar(), nil
}
