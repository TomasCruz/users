package httphandler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/TomasCruz/users/internal/entities"
	"github.com/TomasCruz/users/utils"
	"github.com/TomasCruz/users/utils/errlog"
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
		errlog.Error(err, "HTTPHandler.HealthHandler")
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
		errlog.Error(err, "HTTPHandler.GetUserHandler")
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

// ListUserHandler godoc
// @Summary list users
// @Description list user details
// @ID list-user
// @Produce json
// @Param			country			query			string					false		"Country"
// @Param			page-size		query			string					false		"Page size"
// @Param			page-number		query			string					false		"Page number"
// @Success			200 			{array}			entities.UserResp					"User detail list"
// @Failure			400				{object}		entities.ErrResp					"Bad request"
// @Failure			424				{object}		entities.ErrResp					"Database Error"
// @Failure			500				{object}		entities.ErrResp					"Internal server error"
// @Router /users [get]
func (h HTTPHandler) ListUserHandler(c echo.Context) error {
	values := c.QueryParams()
	filter := utils.ExtractFilter(values)
	userFilter := entities.ExtractUserFilter(filter)
	pageSize, pageNumber := utils.ExtractPagination(values, nil, nil)

	users, totalCount, err := h.cr.ListUser(userFilter, pageSize, pageNumber)
	if err != nil {
		errlog.Error(err, "")
		switch {
		case errors.Is(err, entities.ErrListUser):
			return errorResponse(c, http.StatusFailedDependency, err, entities.ErrListUser.Error())
		case errors.Is(err, entities.ErrCountFilteredQuery):
			return errorResponse(c, http.StatusFailedDependency, err, entities.ErrDatabaseError.Error())
		}

		return errorResponse(c, http.StatusInternalServerError, err, "")
	}

	l := len(users)
	resps := make([]entities.UserResp, 0, l)
	for _, u := range users {
		resps = append(resps, entities.UserRespFromUser(u))
	}

	c.Response().Header().Set("X-Total-Count", strconv.FormatInt(totalCount, 10))
	c.Response().Header().Set("X-Result-Count", strconv.FormatInt(int64(len(users)), 10))
	return c.JSON(http.StatusOK, resps)
}
