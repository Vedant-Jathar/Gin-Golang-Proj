package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger() *zap.Logger{
	// ------------------------------
	// encoders (how logs look)
	// ------------------------------
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		CallerKey:      "caller",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// ------------------------------
	// output destinations
	// ------------------------------
	combinedFile, _ := os.OpenFile("internal/utils/logger/combined.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	errorFile, _ := os.OpenFile("internal/utils/logger/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	combinedWS := zapcore.AddSync(combinedFile)
	errorWS := zapcore.AddSync(errorFile)

	// ------------------------------
	// define log levels
	// ------------------------------
	combinedLevel := zapcore.InfoLevel       // info & above
	errorLevel := zapcore.ErrorLevel         // error & above

	// ------------------------------
	// tee the cores (send log to both)
	// ------------------------------
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, combinedWS, combinedLevel),
		zapcore.NewCore(encoder, errorWS, errorLevel),
	)

	// add option to log caller line number
	Log = zap.New(core, zap.AddCaller())

	return Log
}
