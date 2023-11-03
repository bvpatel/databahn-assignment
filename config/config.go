// config/config.go
package config

import (
	"os"
)

type Config struct {
	KafkaBootstrapServers string
	KafkaTopic            string
	MaxLimit              int64
}

func NewConfig() *Config {
	return &Config{
		KafkaBootstrapServers: os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		KafkaTopic:            os.Getenv("KAFKA_TOPIC"),
		MaxLimit:              1,
	}
}
