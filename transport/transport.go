// Package transport is an interface for synchronous communication
package transport

import (
	"time"
)

type Message struct {
	Header map[string]string
	Body   []byte
}

type Socket interface {
	Recv(*Message) error
	Send(*Message) error
	Close() error
	Local() string
	Remote() string
}

type Client interface {
	Socket
}

type Listener interface {
	Addr() string
	Close() error
	Accept(func(Socket)) error
}

// Transport is an interface which is used for communication between
// services. It uses socket send/recv semantics and had various
// implementations {HTTP, RabbitMQ, NATS, ...}
type Transport interface {
	Init(...Option) error
	Options() Options
	Dial(addr string, opts ...DialOption) (Client, error)
	Listen(addr string, opts ...ListenOption) (Listener, error)
	String() string
}

type Option func(*Options)

type DialOption func(*DialOptions)

type ListenOption func(*ListenOptions)

var (
	DefaultTransport Transport

	DefaultDialTimeout = time.Second * 5

	DefaultTransports   = map[string]func(...Option) Transport{}
	DefaultTransportKey = ""
)

func InitDefaultTransport() {
	if DefaultTransport == nil {
		DefaultTransport = NewTransport()
	}
}

func NewTransport(opts ...Option) Transport {
	if v, ok := DefaultTransports[DefaultTransportKey]; ok {
		return v(opts...)
	}
	panic("DefaultTransportKey is nil")
}
