package app

import (
	l2 "log"
	"os"
	"os/signal"

	"github.com/TomasCruz/users/internal/core/ports"
	"github.com/TomasCruz/users/internal/core/service/app"
	"github.com/TomasCruz/users/internal/handlers/grpchandler"
	"github.com/TomasCruz/users/internal/handlers/httphandler"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/TomasCruz/users/internal/infra/database"
	"github.com/TomasCruz/users/internal/infra/log"
	"github.com/TomasCruz/users/internal/infra/msg"
	"github.com/TomasCruz/users/internal/infra/nts"
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

	// init logger
	logger := log.New(ports.StringToLogLvl(config.MinLogLevel))
	logger.Debug(nil, config.String())
	a.Config = config

	// init DB
	db, err := database.InitDB(config, logger)
	if err != nil {
		logger.Fatal(err, "failed to initialize database")
	}

	// Kafka producer
	msgProducer, err := msg.InitProducer(config, logger)
	if err != nil {
		logger.Fatal(err, "failed to create Kafka producer")
	}

	// NATS
	nc, err := nts.InitNatsConsumer(config, logger)
	if err != nil {
		logger.Fatal(err, "failed to create NATS consumer")
	}
	err = nc.SubUserCreated()
	if err != nil {
		logger.Fatal(err, "failed to subscribe on NATS user created subject")
	}

	// new Service
	svc := app.NewAppUserService(db, msgProducer, nc, logger)

	// init HTTP handler
	e := echo.New()
	h := httphandler.New(e, config.Port, svc, logger)

	// init gRPC handler
	g := grpchandler.New(config.GRPCPort, svc, logger)

	// notify about readiness
	if a.ServerReady != nil {
		a.ServerReady <- struct{}{}
	}

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	gracefulShutdown(db, msgProducer, nc, h, g, logger)
}

func gracefulShutdown(db ports.DB, msgProducer ports.MsgProducer, nc ports.NatsConsumer, h httphandler.HTTPHandler, g *grpchandler.GRPCHandler, logger ports.Logger) {
	// NATS
	err := nc.Unsubscribe()
	if err != nil {
		logger.Error(err, "NATS Unsubscribe failed")
	}

	// Echo
	err = h.Close()
	if err != nil {
		logger.Error(err, "Echo Close failed")
	}

	// gRPC
	g.Close()

	// Kafka
	msgProducer.Close()

	// DB
	err = db.Close()
	if err != nil {
		logger.Error(err, "DB Close failed")
	}

	logger.Debug(nil, "app gracefulShutdown complete")
}
