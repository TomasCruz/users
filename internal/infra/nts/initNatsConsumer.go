package nts

import (
	"github.com/TomasCruz/users/internal/core/ports"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/nats-io/nats.go"
)

func InitNatsConsumer(config configuration.Config, logger ports.Logger) (ports.NatsConsumer, error) {
	nc, err := nats.Connect(config.NatsURL)
	if err != nil {
		return nil, err
	}

	return &natsConsumer{
		nc:     nc,
		config: config,
		logger: logger,
	}, nil
}

type natsConsumer struct {
	nc     *nats.Conn
	sub    *nats.Subscription
	config configuration.Config
	logger ports.Logger
}
