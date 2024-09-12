package httphandler

import (
	"fmt"
	"net/http"

	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/TomasCruz/users/internal/core/service"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	e      *echo.Echo
	svc    service.Service
	logger entities.Logger
}

func New(e *echo.Echo, svc service.Service, port string, logger entities.Logger) HTTPHandler {
	httpHandler := HTTPHandler{e: e, svc: svc, logger: logger}
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
