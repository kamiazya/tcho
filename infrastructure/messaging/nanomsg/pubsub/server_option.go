package pubsub

import (
	"bitbucket.org/kamiazya/tcho/application/event/pubsub"
	"github.com/go-mangos/mangos/protocol/pub"
	"github.com/go-mangos/mangos/transport/all"
	"github.com/go-mangos/mangos/transport/inproc"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	"github.com/go-mangos/mangos/transport/tlstcp"
	"github.com/go-mangos/mangos/transport/ws"
	"github.com/go-mangos/mangos/transport/wss"
)

var (
	DefaultEmmitBufferSize uint = 10
	DefaultURL                  = "ws://127.0.0.1:40899"
)

type Config struct {
	URL             string
	EmmitBufferSize uint

	// protocol
	EnableAll    bool
	EnableICP    bool
	EnableTCP    bool
	EnableINPROC bool
	EnableTLSTCP bool
	EnableWS     bool
	EnableWSS    bool
}

func URL(url string) ServerOption {
	return func(c *Config) error {
		if c.URL != "" {
			return nil
		}
		c.URL = url
		return nil
	}
}

func EmmitBufferSize(emmitBufferSize uint) ServerOption {
	return func(c *Config) error {
		if c.EmmitBufferSize != 0 {
			return nil
		}
		c.EmmitBufferSize = emmitBufferSize
		return nil
	}
}

func EnableAll() ServerOption {
	return func(c *Config) error {
		if c.EnableAll {
			return nil
		}
		c.EnableAll = true
		return nil
	}
}

func EnableICP() ServerOption {
	return func(c *Config) error {
		if c.EnableICP {
			return nil
		}
		c.EnableICP = true
		return nil
	}
}

func EnableTCP() ServerOption {
	return func(c *Config) error {
		if c.EnableTCP {
			return nil
		}
		c.EnableTCP = true
		return nil
	}
}

func EnableINPROC() ServerOption {
	return func(c *Config) error {
		if c.EnableINPROC {
			return nil
		}
		c.EnableINPROC = true
		return nil
	}
}

func EnableTLSTCP() ServerOption {
	return func(c *Config) error {
		if c.EnableTLSTCP {
			return nil
		}
		c.EnableTLSTCP = true
		return nil
	}
}
func EnableWS() ServerOption {
	return func(c *Config) error {
		if c.EnableWS {
			return nil
		}
		c.EnableWS = true
		return nil
	}
}
func EnableWSS() ServerOption {
	return func(c *Config) error {
		if c.EnableWSS {
			return nil
		}
		c.EnableWSS = true
		return nil
	}
}

type ServerOption func(*Config) error

func genFromConfig(c *Config) (pubsub.Server, error) {
	if c.URL == "" {
		c.URL = DefaultURL
	}

	if c.EnableAll == false && c.EnableICP == false && c.EnableTCP == false && c.EnableINPROC == false && c.EnableTLSTCP == false && c.EnableWS == false && c.EnableWSS == false {
		c.EnableWS = true
	}
	if c.EmmitBufferSize > 0 {
		c.EmmitBufferSize = DefaultEmmitBufferSize
	}

	var (
		err error
		s   = &nanomsgServer{
			ch: make(chan pubsub.Event, c.EmmitBufferSize),
		}
	)
	if s.sock, err = pub.NewSocket(); err != nil {
		return nil, err
	}
	if c.EnableAll {
		all.AddTransports(s.sock)
	} else {
		if c.EnableICP {
			s.sock.AddTransport(ipc.NewTransport())
		}
		if c.EnableTCP {
			s.sock.AddTransport(tcp.NewTransport())
		}
		if c.EnableINPROC {
			s.sock.AddTransport(inproc.NewTransport())
		}
		if c.EnableTLSTCP {
			s.sock.AddTransport(tlstcp.NewTransport())
		}
		if c.EnableWS {
			s.sock.AddTransport(ws.NewTransport())
		}
		if c.EnableWSS {
			s.sock.AddTransport(wss.NewTransport())
		}
	}

	if err = s.sock.Listen(s.url); err != nil {
		return nil, err
	}

	return s, nil
}
