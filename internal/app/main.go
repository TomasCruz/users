package app

import (
	"os"
	"os/signal"

	"github.com/TomasCruz/users/internal/configuration"
	"github.com/TomasCruz/users/internal/core"
	"github.com/TomasCruz/users/internal/database"
	"github.com/TomasCruz/users/internal/entities"
	"github.com/TomasCruz/users/internal/handlers/httphandler"
	"github.com/TomasCruz/users/internal/msg"
	"github.com/labstack/echo/v4"
)

type App struct {
	EnvFile     string
	Config      configuration.Config
	ServerReady chan struct{}
}

func (application *App) Start() {
	// populate configuration
	if application.EnvFile == "" {
		application.EnvFile = ".env"
	}

	config, err := ConfigFromEnvVars(application.EnvFile)
	if err != nil {
		entities.LogFatal(err, "failed to read environment variables")
	}
	entities.LogDebug(nil, config.String())
	application.Config = config

	// init DB
	db, err := database.InitDB(config)
	if err != nil {
		entities.LogFatal(err, "failed to initialize database")
	}

	// Kafka producer
	msg, err := msg.InitMsg(config)
	if err != nil {
		entities.LogFatal(err, "failed to create Kafka producer")
	}

	// new Core
	cr := core.New(config, db, msg)

	// init HTTP handler
	e := echo.New()
	h := httphandler.New(e, cr, config)

	// notify about readiness
	if application.ServerReady != nil {
		application.ServerReady <- struct{}{}
	}

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	gracefulShutdown(db, msg, h)
}

func gracefulShutdown(db core.DB, msg core.Msg, h httphandler.HTTPHandler) {
	// Echo
	err := h.Close()
	if err != nil {
		entities.LogError(err, "Echo Close failed")
	}

	// Kafka
	msg.Close()

	// DB
	err = db.Close()
	if err != nil {
		entities.LogError(err, "DB Close failed")
	}
}
