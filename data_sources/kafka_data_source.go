package data_sources

import (
	"databahn-api/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaDataSource struct {
	kafkaProducer *kafka.Producer
}

func NewKafkaDataSource(config *config.Config) (*KafkaDataSource, error) {
	kafkaProducer, err := createKafkaProducer(config)
	if err != nil {
		return nil, err
	}
	return &KafkaDataSource{
		kafkaProducer: kafkaProducer,
	}, nil
}

func createKafkaProducer(config *config.Config) (*kafka.Producer, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": config.KafkaBootstrapServers,
	}
	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func (ds *KafkaDataSource) PushData(data string) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &config.KafkaTopic, Partition: kafka.PartitionAny},
		Value:          []byte(data),
	}
	if err := ds.kafkaProducer.Produce(message, nil); err != nil {
		return err
	}
	return nil
}
