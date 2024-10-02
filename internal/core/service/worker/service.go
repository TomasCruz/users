package worker

import (
	"github.com/TomasCruz/users/internal/core/ports"
)

type WorkerUserService struct {
	msgProducer ports.MsgProducer
	logger      ports.Logger
}

func NewWorkerUserService(msgProducer ports.MsgProducer, logger ports.Logger) WorkerUserService {
	return WorkerUserService{
		msgProducer: msgProducer,
		logger:      logger,
	}
}
