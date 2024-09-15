package msg

import (
	"encoding/json"

	"github.com/TomasCruz/users/internal/domain/entities"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (k *kafkaMsgConsumer) consume() {
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
				err := k.consumeUserMsg(topic, eType.Value)
				if err != nil {
					k.logger.Error(err, "ConsumeUserCreatedMsg")
				}
			}
		case kafka.Error:
			// Handle Kafka errors
			k.logger.Error(nil, ev.String())
		}
	}

	k.kc.Close()
	k.shutdownComplete <- struct{}{}
}

func (k *kafkaMsgConsumer) consumeUserMsg(topic string, userBytes []byte) error {
	var user entities.User
	err := json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}

	switch topic {
	case k.config.CreateUserTopic:
		if err = k.cr.ConsumeUserCreatedMsg(user); err != nil {
			return err
		}
	}

	return nil
}
