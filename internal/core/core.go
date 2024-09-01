package core

import (
	"github.com/TomasCruz/users/internal/configuration"
)

type Core struct {
	config configuration.Config
	db     DB
	kp     Msg
}

func New(config configuration.Config, db DB, msg Msg) Core {
	return Core{
		config: config,
		db:     db,
		kp:     msg,
	}
}
