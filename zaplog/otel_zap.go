package zaplog

import (
	"os"

	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel/log/noop"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// not good practice to have global variable, but for simplicity, we will use it here.
var Logger *zap.Logger

func InitLogger() {
	provider := noop.NewLoggerProvider()

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
		otelzap.NewCore("sre-works", otelzap.WithLoggerProvider(provider)),
	)

	Logger = zap.New(core)
}
