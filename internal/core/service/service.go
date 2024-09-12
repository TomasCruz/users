package core

import "github.com/TomasCruz/users/internal/core/entities"

type Service struct {
	db     entities.DB
	msg    entities.Msg
	logger entities.Logger
}

func New(db entities.DB, msg entities.Msg, logger entities.Logger) Service {
	return Service{
		db:     db,
		msg:    msg,
		logger: logger,
	}
}
