package core

import (
	"github.com/TomasCruz/users/internal/domain/ports"
)

type Core struct {
	db     ports.DB
	msg    ports.Msg
	logger ports.Logger
}

func New(db ports.DB, msg ports.Msg, logger ports.Logger) Core {
	return Core{
		db:     db,
		msg:    msg,
		logger: logger,
	}
}
