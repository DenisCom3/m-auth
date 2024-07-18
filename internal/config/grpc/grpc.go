package grpc

import (
	"fmt"
	"net"
)

type Grpc struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (g Grpc) Address() string {
	return net.JoinHostPort(g.Host, fmt.Sprintf("%v", g.Port))
}
