package redis

import (
	"fmt"
	"net"
	"time"
)

type Redis struct {
	Host             string        `yaml:"host"`
	Port             int           `yaml:"port"`
	ConnTimeout      time.Duration `yaml:"connection-timeout" env-default:"5s"`
	MaxIdleValue     int           `yaml:"max-idle"`
	IdleTimeoutValue time.Duration `yaml:"idle-timeout"`
}

func (r Redis) Address() string {
	return net.JoinHostPort(r.Host, fmt.Sprintf("%v", r.Port))
}

func (r Redis) ConnectionTimeout() time.Duration {
	return r.ConnTimeout
}

func (r Redis) MaxIdle() int {
	return r.MaxIdleValue
}

func (r Redis) IdleTimeout() time.Duration {
	return r.IdleTimeoutValue
}
