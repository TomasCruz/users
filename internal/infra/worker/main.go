package worker

import (
	l2 "log"
	"os"
	"os/signal"

	"github.com/TomasCruz/users/internal/infra/log"
	"github.com/TomasCruz/users/internal/infra/msg"

	"github.com/TomasCruz/users/internal/domain/core"
	"github.com/TomasCruz/users/internal/domain/ports"
	"github.com/TomasCruz/users/internal/infra/configuration"
)

type WorkerApp struct {
	EnvFile string
	Config  configuration.Config
	// kc      *kafka.Consumer
}

func (w *WorkerApp) Start() {
	// populate configuration
	if w.EnvFile == "" {
		w.EnvFile = ".env"
	}

	config, err := configuration.ConfigFromEnvVars(w.EnvFile)
	if err != nil {
		l2.Fatal("failed to read environment variables", err)
	}

	// init logger
	logger := log.New(ports.StringToLogLvl(config.MinLogLevel))
	logger.Debug(nil, config.String())
	w.Config = config

	// new Service
	cr := core.New(nil, nil, logger)

	// Kafka consumer
	msgConsumer, err := msg.InitConsumer(config, cr, logger)
	if err != nil {
		logger.Fatal(err, "failed to create Kafka producer")
	}

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	gracefulShutdown(msgConsumer, logger)
}

func gracefulShutdown(msgConsumer ports.MsgConsumer, logger ports.Logger) {
	// Kafka
	msgConsumer.Close()

	logger.Info(nil, "worker gracefulShutdown complete")
}
