package app

import (
	"github.com/TomasCruz/users/internal/core/ports"
)

type AppUserService struct {
	db           ports.DB
	msgProducer  ports.MsgProducer
	natsConsumer ports.NatsConsumer
	logger       ports.Logger
}

func NewAppUserService(db ports.DB, msgProducer ports.MsgProducer, natsConsumer ports.NatsConsumer, logger ports.Logger) AppUserService {
	return AppUserService{
		db:           db,
		msgProducer:  msgProducer,
		natsConsumer: natsConsumer,
		logger:       logger,
	}
}
