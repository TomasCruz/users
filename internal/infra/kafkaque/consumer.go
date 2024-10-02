package kafkaque

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (k *kafkaConsumer) consume() {
	for !k.shutdownReceived {
		// Poll for Kafka messages
		ev := k.kc.Poll(100)
		if ev == nil {
			continue
		}

		switch eType := ev.(type) {
		case *kafka.Message:
			// Process the consumed message
			topic := *eType.TopicPartition.Topic
			switch topic {
			case k.config.CreateUserTopic:
				err := k.consumeUserCreatedEvent(eType.Value)
				if err != nil {
					k.logger.Error(err, "consumeUserCreatedMsg")
				}
			}
		case kafka.Error:
			// Handle Kafka errors
			k.logger.Error(nil, eType.String())
		}
	}

	k.kc.Close()
	k.shutdownComplete <- struct{}{}
}
