package httphandler

import (
	"fmt"
	"net/http"

	"github.com/TomasCruz/users/internal/core/ports"
	"github.com/TomasCruz/users/internal/core/service/app"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	e      *echo.Echo
	svc    app.AppUserService
	logger ports.Logger
}

func New(e *echo.Echo, port string, svc app.AppUserService, logger ports.Logger) HTTPHandler {
	httpHandler := HTTPHandler{e: e, svc: svc, logger: logger}
	httpHandler.bindRoutes()

	// fire up the HTTP server
	go func() {
		err := e.Start(fmt.Sprintf(":%s", port))
		if err != nil && err != http.ErrServerClosed {
			logger.Error(err, "Echo error")
		}
	}()

	return httpHandler
}
