package httphandler

import (
	"errors"
	"net/http"

	"github.com/TomasCruz/users/internal/entities"
	"github.com/TomasCruz/users/internal/errstack"
	"github.com/labstack/echo/v4"
)

// Health godoc
// @Summary health check
// @ID health
// @Produce	json
// @Success			204 																"Healthy"
// @Failure			500				{object}		entities.ErrResp					"Internal server error"
// @Router /health [get]
func (h HTTPHandler) HealthHandler(c echo.Context) error {
	err := h.cr.Health()
	if err != nil {
		errstack.Error(err, "HTTPHandler.HealthHandler")
		return errorResponse(c, http.StatusInternalServerError, err, "")
	}

	return c.NoContent(http.StatusNoContent)
}

// GetUserHandler godoc
// @Summary gets user
// @Description gets user details
// @ID get-user
// @Produce json
// @Param			user-id			path			string					true		"User id"
// @Success			200 			{object}		entities.UserResp					"User details"
// @Failure			400				{object}		entities.ErrResp					"Bad ID"
// @Failure			404				{object}		entities.ErrResp					"Not found"
// @Failure			424				{object}		entities.ErrResp					"Database Error"
// @Failure			500				{object}		entities.ErrResp					"Internal server error"
// @Router /users/{user-id} [get]
func (h HTTPHandler) GetUserHandler(c echo.Context) error {
	uuidString := c.Param("user-id")
	userID, err := h.validateUUID(uuidString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrResp{Msg: err.Error()})
	}

	user, err := h.cr.GetUserByID(userID)
	if err != nil {
		errstack.Error(err, "HTTPHandler.GetUserHandler")
		switch {
		case errors.Is(err, entities.ErrNonexistingUser):
			return errorResponse(c, http.StatusNotFound, err, entities.ErrNonexistingUser.Error())
		case errors.Is(err, entities.ErrGetUser):
			return errorResponse(c, http.StatusFailedDependency, err, entities.ErrGetUser.Error())
		}

		return errorResponse(c, http.StatusInternalServerError, err, "")
	}

	resp := entities.UserRespFromUser(user)
	return c.JSON(http.StatusOK, resp)
}
