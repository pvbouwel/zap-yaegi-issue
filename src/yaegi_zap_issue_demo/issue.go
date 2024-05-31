package yaegi_zap_issue_demo

import (
	"context"
	"net/http"
)

type Config struct {
	LogLevel            string `yaml:"loglevel"`
}

func CreateConfig() *Config {
	return &Config{}
}

type issuePlugin struct {
	name                string
	next                http.Handler
}

func (p *issuePlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	p.next.ServeHTTP(rw, req)
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &issuePlugin{
		name:                name,
		next:                next,
	}, nil
}
