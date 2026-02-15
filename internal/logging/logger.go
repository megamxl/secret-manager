package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Loggers struct {
	App   *zap.Logger // Application logs
	Audit *zap.Logger // Audit trail
}

func Setup(level, format, auditFile string) (*Loggers, error) {
	// Application logger
	appLogger := createLogger(level, format, os.Stdout)

	// Audit logger (always JSON, separate file)
	auditLogger := createAuditLogger(auditFile)

	return &Loggers{
		App:   appLogger,
		Audit: auditLogger,
	}, nil
}

func createLogger(level, format string, output *os.File) *zap.Logger {
	lvl, _ := zapcore.ParseLevel(level)

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "time",
		LevelKey:    "level",
		MessageKey:  "msg",
		EncodeTime:  zapcore.ISO8601TimeEncoder,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	}

	var encoder zapcore.Encoder
	if format == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(output), lvl)
	return zap.New(core)
}

func createAuditLogger(path string) *zap.Logger {
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:    "@timestamp",
		MessageKey: "action",
		EncodeTime: zapcore.RFC3339NanoTimeEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(file),
		zapcore.InfoLevel,
	)

	return zap.New(core, zap.Fields(zap.String("category", "audit")))
}

func (l *Loggers) Close() {
	l.App.Sync()
	l.Audit.Sync()
}
