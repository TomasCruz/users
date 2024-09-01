package httphandler

import (
	"net/http"

	"github.com/TomasCruz/users/internal/core"
	"github.com/labstack/echo/v4"
)

// Health godoc
// @Summary health check
// @Description display health status
// @ID health
// @Produce	json
// @Success			204 																"Healthy"
// @Failure			500				{object}		model.ErrResp						"Internal server error"
// @Router /health [get]
func (h HTTPHandler) HealthHandler(c echo.Context) error {
	cr := core.New(h.config, h.db, h.msg)
	err := cr.Health()
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, err, "")
	}

	return c.NoContent(http.StatusNoContent)
}
