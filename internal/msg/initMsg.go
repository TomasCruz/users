package msg

import (
	"log"

	"github.com/TomasCruz/users/internal/configuration"
	"github.com/TomasCruz/users/internal/core"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func InitMsg(config configuration.Config) (core.Msg, error) {
	var kp *kafka.Producer
	var err error
	if kp, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaBroker,
	}); err != nil {
		log.Fatalf("failed to create Kafka producer: %s", err.Error())
	}

	return kafkaMsg{kp: kp}, nil
}

type kafkaMsg struct {
	kp *kafka.Producer
}
