package worker

import (
	"github.com/TomasCruz/users/internal/core/ports"
)

type WorkerUserService struct {
	logger ports.Logger
}

func NewWorkerUserService(logger ports.Logger) WorkerUserService {
	return WorkerUserService{
		logger: logger,
	}
}
