package app

import (
	"github.com/TomasCruz/users/internal/core/ports"
)

type AppUserService struct {
	db            ports.DB
	queueProducer ports.QueueProducer
	msgConsumer   ports.MsgConsumer
	logger        ports.Logger
}

func NewAppUserService(db ports.DB, queueProducer ports.QueueProducer, msgConsumer ports.MsgConsumer, logger ports.Logger) AppUserService {
	return AppUserService{
		db:            db,
		queueProducer: queueProducer,
		msgConsumer:   msgConsumer,
		logger:        logger,
	}
}
