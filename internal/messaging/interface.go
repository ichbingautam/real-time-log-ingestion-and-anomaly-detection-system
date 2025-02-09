package messaging

import "context"

import "time"

// Add missing time import and config references
type TLSConfig struct {
	Enabled   bool
	CAFile    string
	CertFile  string
	KeyFile   string
}

type SASLConfig struct {
	Enabled  bool
	Username string
	Password string
}

type SecurityConfig struct {
	TLS  TLSConfig
	SASL SASLConfig
}

// Ingester defines the interface for message consumers
type Ingester interface {
	Consume(ctx context.Context, handler func([]byte) error) error
	Commit() error
	Close() error
}

// Egester defines the interface for message producers
type Egester interface {
	Produce(ctx context.Context, messages [][]byte) error
	Close() error
}

// Message represents a generic message structure
type Message struct {
	Topic     string
	Key       []byte
	Value     []byte
	Timestamp time.Time
	Headers   map[string]string
}

// Handler defines the message processing function signature
type Handler func(Message) error