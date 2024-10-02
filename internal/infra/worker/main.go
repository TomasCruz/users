package worker

import (
	l2 "log"
	"os"
	"os/signal"

	"github.com/TomasCruz/users/internal/infra/kafkaque"
	"github.com/TomasCruz/users/internal/infra/log"
	"github.com/TomasCruz/users/internal/infra/natsmsg"

	"github.com/TomasCruz/users/internal/core/ports"
	"github.com/TomasCruz/users/internal/core/service/worker"
	"github.com/TomasCruz/users/internal/infra/configuration"
)

type WorkerApp struct {
	EnvFile string
	Config  configuration.Config
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

	// NATS
	np, err := natsmsg.InitProducer(config, logger)
	if err != nil {
		logger.Fatal(err, "failed to create NATS producer")
	}

	// new Service
	svc := worker.NewWorkerUserService(np, logger)

	// Kafka consumer
	kafkaConsumer, err := kafkaque.InitConsumer(config, svc, logger)
	if err != nil {
		logger.Fatal(err, "failed to create Kafka consumer")
	}

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	gracefulShutdown(kafkaConsumer, np, logger)
}

func gracefulShutdown(kafkaConsumer ports.QueueConsumer, np ports.MsgProducer, logger ports.Logger) {
	// NATS
	err := np.Close()
	if err != nil {
		logger.Error(err, "NATS Drain failed")
	}

	// Kafka
	kafkaConsumer.Close()

	logger.Debug(nil, "worker gracefulShutdown complete")
}
