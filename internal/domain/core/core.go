package core

import "github.com/TomasCruz/users/internal/domain/entities"

type Core struct {
	db     entities.DB
	msg    entities.Msg
	logger entities.Logger
}

func New(db entities.DB, msg entities.Msg, logger entities.Logger) Core {
	return Core{
		db:     db,
		msg:    msg,
		logger: logger,
	}
}
