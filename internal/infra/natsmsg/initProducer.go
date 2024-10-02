package natsmsg

import (
	"github.com/TomasCruz/users/internal/core/ports"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/nats-io/nats.go"
)

func InitProducer(config configuration.Config, logger ports.Logger) (ports.MsgProducer, error) {
	nc, err := nats.Connect(config.NatsURL)
	if err != nil {
		return nil, err
	}

	return natsProducer{
		nc:     nc,
		config: config,
		logger: logger,
	}, nil
}

type natsProducer struct {
	nc     *nats.Conn
	config configuration.Config
	logger ports.Logger
}
