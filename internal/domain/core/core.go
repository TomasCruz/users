package core

import (
	"github.com/TomasCruz/users/internal/domain/ports"
)

type Core struct {
	db          ports.DB
	msgProducer ports.MsgProducer
	logger      ports.Logger
}

func New(db ports.DB, msgProducer ports.MsgProducer, logger ports.Logger) Core {
	return Core{
		db:          db,
		msgProducer: msgProducer,
		logger:      logger,
	}
}
