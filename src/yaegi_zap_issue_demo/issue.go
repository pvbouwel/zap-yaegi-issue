package yaegi_zap_issue_demo

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel            string `yaml:"loglevel"`
}

func CreateConfig() *Config {
	return &Config{}
}

type issuePlugin struct {
	name   string
	next   http.Handler
	logger *zap.Logger
}

func (p *issuePlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	p.logger.Info("test")
	p.next.ServeHTTP(rw, req)
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	logger, _ := NewLogger("INFO")
	return &issuePlugin{
		name:   name,
		next:   next,
		logger: logger,
	}, nil
}

func NewLogger(logLevel string) (*zap.Logger, error) {
	config := zap.Config{
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey: "time",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z"))
			},
			LevelKey:      "level",
			NameKey:       "logger",
			MessageKey:    "message",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	var level zapcore.Level
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		panic(err)
	}
	config.Level = zap.NewAtomicLevelAt(level)
	return config.Build()
}