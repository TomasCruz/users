package httphandler

import (
	"fmt"
	"net/http"

	"github.com/TomasCruz/users/internal/domain/core"
	"github.com/TomasCruz/users/internal/domain/ports"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	e      *echo.Echo
	cr     core.Core
	logger ports.Logger
}

func New(e *echo.Echo, cr core.Core, port string, logger ports.Logger) HTTPHandler {
	httpHandler := HTTPHandler{e: e, cr: cr, logger: logger}
	httpHandler.bindRoutes()

	// fire up the server
	go func() {
		err := e.Start(fmt.Sprintf(":%s", port))
		if err != nil && err != http.ErrServerClosed {
			logger.Error(err, "Echo error")
		}
	}()

	return httpHandler
}
