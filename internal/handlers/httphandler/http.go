package httphandler

import (
	"fmt"
	"net/http"

	"github.com/TomasCruz/users/internal/configuration"
	"github.com/TomasCruz/users/internal/core"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	e      *echo.Echo
	cr     core.Core
	config configuration.Config
}

func New(e *echo.Echo, cr core.Core, config configuration.Config) HTTPHandler {
	httpHandler := HTTPHandler{e: e, cr: cr, config: config}
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
