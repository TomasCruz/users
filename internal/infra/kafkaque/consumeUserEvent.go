package kafkaque

import (
	"encoding/json"

	"github.com/TomasCruz/users/internal/core/entities"
)

func (k *kafkaConsumer) consumeUserCreatedEvent(userBytes []byte) error {
	var user entities.User
	err := json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}

	if err = k.svc.ConsumeUserCreatedEvent(user); err != nil {
		return err
	}

	return nil
}
