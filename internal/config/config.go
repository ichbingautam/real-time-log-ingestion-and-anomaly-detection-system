package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Kafka      KafkaConfig      `yaml:"kafka"`
	Elastic    ElasticConfig    `yaml:"elastic"`
	S3         S3Config         `yaml:"s3"`
	Monitoring MonitoringConfig `yaml:"monitoring"`
}

type KafkaConfig struct {
	Brokers  []string `yaml:"brokers"`
	Topic    string   `yaml:"topic"`
	GroupID  string   `yaml:"group_id"`
	Security SecurityConfig `yaml:"security"`
}

type SecurityConfig struct {
	TLS  TLSConfig  `yaml:"tls"`
	SASL SASLConfig `yaml:"sasl"`
}

type TLSConfig struct {
	Enabled bool   `yaml:"enabled"`
	CAFile  string `yaml:"ca_file"`
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
}

type SASLConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Add missing type definitions
type ElasticConfig struct {
	Addresses []string `yaml:"addresses"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
	Index     string   `yaml:"index"`
}

type S3Config struct {
	Region    string `yaml:"region"`
	Bucket    string `yaml:"bucket"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}

type MonitoringConfig struct {
	Prometheus struct {
		Enabled bool `yaml:"enabled"`
		Port    int  `yaml:"port"`
	} `yaml:"prometheus"`
}