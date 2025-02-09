package messaging

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/config"
	"github.com/segmentio/kafka-go" // Now used in struct definitions
	"github.com/segmentio/kafka-go/sasl/plain"
)

// Add concrete implementation using context and kafka package
type KafkaClient struct {
	reader *kafka.Reader
	writer *kafka.Writer
}

func NewKafkaClient(cfg config.Config) *KafkaClient {
	dialer := createDialer(cfg.Security)

	return &KafkaClient{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: cfg.Brokers,
			Topic:   cfg.Topic,
			GroupID: cfg.GroupID,
			Dialer:  dialer,
		}),
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: cfg.Brokers,
			Topic:   cfg.Topic,
			Dialer:  dialer,
		}),
	}
}

// Add concrete TLS/SASL config implementations
func createDialer(cfg SecurityConfig) *kafka.Dialer {
	tlsConfig, err := createTLSConfig(cfg.TLS)
	if err != nil {
		panic(err)
	}

	saslMechanism := createSASLConfig(cfg.SASL)

	return &kafka.Dialer{
		TLS:           tlsConfig,
		SASLMechanism: saslMechanism,
	}
}

func createTLSConfig(cfg TLSConfig) (*tls.Config, error) {
	if !cfg.Enabled {
		return nil, nil
	}

	caCert, err := os.ReadFile(cfg.CAFile)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}, nil
}

func createSASLConfig(cfg SASLConfig) plain.Mechanism {
	if !cfg.Enabled {
		return plain.Mechanism{}
	}

	return plain.Mechanism{
		Username: cfg.Username,
		Password: cfg.Password,
	}
}