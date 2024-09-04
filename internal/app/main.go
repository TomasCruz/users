package app

import (
	"os"
	"os/signal"

	"github.com/TomasCruz/users/internal/core"
	"github.com/TomasCruz/users/internal/database"
	"github.com/TomasCruz/users/internal/errlog"
	"github.com/TomasCruz/users/internal/handlers/httphandler"
	"github.com/TomasCruz/users/internal/msg"
	"github.com/labstack/echo/v4"
)

func Start() {
	// populate configuration
	config, err := setupFromEnvVars()
	if err != nil {
		errlog.Fatal(err, "failed to read environment variables")
	}

	// init DB
	db, err := database.InitDB(config)
	if err != nil {
		errlog.Fatal(err, "failed to initialize database")
	}

	// Kafka producer
	msg, err := msg.InitMsg(config)
	if err != nil {
		errlog.Fatal(err, "failed to create Kafka producer")
	}

	// new Core
	cr := core.New(config, db, msg)

	// init HTTP handler
	e := echo.New()
	h := httphandler.New(e, cr, config)

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
		errlog.Error(err, "Echo Close failed")
	}

	// Kafka
	msg.Close()

	// DB
	err = db.Close()
	if err != nil {
		errlog.Error(err, "DB Close failed")
	}
}
