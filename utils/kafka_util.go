package util

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaBatchProducer struct {
	producer *kafka.Producer
}

func NewKafkaBatchProducer(bootstrapServers string) (*KafkaBatchProducer, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	}
	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		return nil, err
	}
	return &KafkaBatchProducer{producer: producer}, nil
}

func (kp *KafkaBatchProducer) ProduceMessages(topic string, messages []string) error {
	for _, message := range messages {
		kafkaMessage := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}
		if err := kp.producer.Produce(kafkaMessage, nil); err != nil {
			return err
		}
	}
	return nil
}

func (kp *KafkaBatchProducer) Close() {
	kp.producer.Close()
}
