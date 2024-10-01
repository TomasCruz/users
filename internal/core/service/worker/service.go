package worker

import (
	"github.com/TomasCruz/users/internal/core/ports"
)

type WorkerUserService struct {
	np     ports.NatsProducer
	logger ports.Logger
}

func NewWorkerUserService(np ports.NatsProducer, logger ports.Logger) WorkerUserService {
	return WorkerUserService{
		np:     np,
		logger: logger,
	}
}
