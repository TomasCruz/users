package natsmsg

import (
	"github.com/TomasCruz/users/internal/core/ports"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/nats-io/nats.go"
)

func InitConsumer(config configuration.Config, logger ports.Logger) (ports.MsgConsumer, error) {
	nc, err := nats.Connect(config.NatsURL)
	if err != nil {
		return nil, err
	}

	nCons := &natsConsumer{
		nc:     nc,
		config: config,
		logger: logger,
	}

	err = nCons.subUserCreated()
	if err != nil {
		return nil, err
	}

	return nCons, nil
}

type natsConsumer struct {
	nc     *nats.Conn
	sub    *nats.Subscription
	config configuration.Config
	logger ports.Logger
}
