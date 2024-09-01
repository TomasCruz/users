package httphandler

import (
	"fmt"
	"net/http"

	"github.com/TomasCruz/users/internal/configuration"
	"github.com/TomasCruz/users/internal/core"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HTTPHandler struct {
	e      *echo.Echo
	db     core.DB
	msg    core.Msg
	config configuration.Config
}

func New(config configuration.Config, db core.DB, msg core.Msg) core.Http {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	httpHandler := HTTPHandler{e: e, db: db, msg: msg, config: config}
	httpHandler.bindRoutes()

	// fire up the server
	go func() {
		err := e.Start(fmt.Sprintf(":%s", config.Port))
		if err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("Exiting: %s", err.Error())
		}
	}()

	return httpHandler
}
