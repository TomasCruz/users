package app

import (
	l2 "log"
	"os"
	"os/signal"

	"github.com/TomasCruz/users/internal/adapters/httphandler"
	"github.com/TomasCruz/users/internal/domain/core"
	"github.com/TomasCruz/users/internal/domain/ports"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/TomasCruz/users/internal/infra/database"
	"github.com/TomasCruz/users/internal/infra/log"
	"github.com/TomasCruz/users/internal/infra/msg"
	"github.com/labstack/echo/v4"
)

type App struct {
	EnvFile     string
	Config      configuration.Config
	ServerReady chan struct{}
}

func (a *App) Start() {
	// populate configuration
	if a.EnvFile == "" {
		a.EnvFile = ".env"
	}

	config, err := configuration.ConfigFromEnvVars(a.EnvFile)
	if err != nil {
		l2.Fatal("failed to read environment variables", err)
	}

	logger := log.New(ports.StringToLogLvl(config.MinLogLevel))
	logger.Debug(nil, config.String())
	a.Config = config

	// init DB
	db, err := database.InitDB(config, logger)
	if err != nil {
		logger.Fatal(err, "failed to initialize database")
	}

	// Kafka producer
	msg, err := msg.InitMsg(config, logger)
	if err != nil {
		logger.Fatal(err, "failed to create Kafka producer")
	}

	// new Service
	s := core.New(db, msg, logger)

	// init HTTP handler
	e := echo.New()
	h := httphandler.New(e, s, config.Port, logger)

	// notify about readiness
	if a.ServerReady != nil {
		a.ServerReady <- struct{}{}
	}

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	gracefulShutdown(db, msg, h, logger)
}

func gracefulShutdown(db ports.DB, msg ports.Msg, h httphandler.HTTPHandler, logger ports.Logger) {
	// Echo
	err := h.Close()
	if err != nil {
		logger.Error(err, "Echo Close failed")
	}

	// Kafka
	msg.Close()

	// DB
	err = db.Close()
	if err != nil {
		logger.Error(err, "DB Close failed")
	}
}
