package httphandler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/TomasCruz/users/internal/domain/entities"
	"github.com/labstack/echo/v4"
)

// Health godoc
// @Summary health check
// @ID health
// @Produce	json
// @Success			204 																"Healthy"
// @Failure			500				{object}		httphandler.ErrResp					"Internal server error"
// @Router /health [get]
func (h HTTPHandler) HealthHandler(c echo.Context) error {
	err := h.cr.Health()
	if err != nil {
		h.logger.Error(err, "HTTPHandler.HealthHandler")
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
// @Success			200 			{object}		httphandler.UserResp				"User details"
// @Failure			400				{object}		httphandler.ErrResp					"Bad ID"
// @Failure			404				{object}		httphandler.ErrResp					"Not found"
// @Failure			424				{object}		httphandler.ErrResp					"Database Error"
// @Failure			500				{object}		httphandler.ErrResp					"Internal server error"
// @Router /users/{user-id} [get]
func (h HTTPHandler) GetUserHandler(c echo.Context) error {
	uuidString := c.Param("user-id")
	userID, err := h.validateUUID(uuidString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrResp{Msg: err.Error()})
	}

	user, err := h.cr.GetUserByID(userID)
	if err != nil {
		h.logger.Error(err, "HTTPHandler.GetUserHandler")

		switch {
		case errors.Is(err, entities.ErrNonexistingUser):
			return errorResponse(c, http.StatusNotFound, err, entities.ErrNonexistingUser.Error())
		case errors.Is(err, entities.ErrGetUser):
			return errorResponse(c, http.StatusFailedDependency, err, entities.ErrGetUser.Error())
		}

		return errorResponse(c, http.StatusInternalServerError, err, "")
	}

	resp := userRespFromUser(user)
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
// @Success			200 			{array}			httphandler.UserResp				"User detail list"
// @Failure			400				{object}		httphandler.ErrResp					"Bad request"
// @Failure			424				{object}		httphandler.ErrResp					"Database Error"
// @Failure			500				{object}		httphandler.ErrResp					"Internal server error"
// @Router /users [get]
func (h HTTPHandler) ListUserHandler(c echo.Context) error {
	values := c.QueryParams()
	filter := extractFilterAndPaginatorFromQueryParams(values)
	userFilter := entities.ExtractUserFilter(filter)
	pageSize, pageNumber := entities.ExtractPagination(filter, nil, nil)

	users, totalCount, err := h.cr.ListUser(userFilter, pageSize, pageNumber)
	if err != nil {
		h.logger.Error(err, "HTTPHandler.ListUserHandler")

		switch {
		case errors.Is(err, entities.ErrListUser):
			return errorResponse(c, http.StatusFailedDependency, err, entities.ErrListUser.Error())
		case errors.Is(err, entities.ErrCountFilteredQuery):
			return errorResponse(c, http.StatusFailedDependency, err, entities.ErrDatabaseError.Error())
		}

		return errorResponse(c, http.StatusInternalServerError, err, "")
	}

	l := len(users)
	resps := make([]UserResp, 0, l)
	for _, u := range users {
		resps = append(resps, userRespFromUser(u))
	}

	c.Response().Header().Set("X-Total-Count", strconv.FormatInt(totalCount, 10))
	c.Response().Header().Set("X-Result-Count", strconv.FormatInt(int64(len(users)), 10))
	return c.JSON(http.StatusOK, resps)
}

// CreateUser godoc
// @Summary creates user
// @Description creates user
// @ID create-user
// @Consume	json
// @Produce	json
// @Param			payload			body			httphandler.CreateUserReq	true	"Payload"
// @Success			201 			{object}		httphandler.UserResp				"User details"
// @Failure			400				{object}		httphandler.ErrResp					"Bad request"
// @Failure			409				{object}		httphandler.ErrResp					"Existing email"
// @Failure			424				{object}		httphandler.ErrResp					"Database or Kafka Error"
// @Failure			500				{object}		httphandler.ErrResp					"Internal server error"
// @Router /users [put]
func (h HTTPHandler) CreateUserHandler(c echo.Context) error {
	req := CreateUserReq{}
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, err, "")
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err, "")
	}

	if err := h.validateCreateUser(req); err != nil {
		return errorResponse(c, http.StatusBadRequest, err, "")
	}

	userDTO := userDTOFromCreateUserReq(req)
	user, err := h.cr.CreateUser(userDTO)
	if err != nil {
		h.logger.Error(err, "HTTPHandler.CreateUserHandler")

		switch {
		case errors.Is(err, entities.ErrBadEmail):
			return errorResponse(c, http.StatusBadRequest, err, entities.ErrBadEmail.Error())
		case errors.Is(err, entities.ErrEmptyName):
			return errorResponse(c, http.StatusBadRequest, err, entities.ErrEmptyName.Error())
		case errors.Is(err, entities.ErrBadHashedPswd):
			return errorResponse(c, http.StatusBadRequest, err, entities.ErrBadHashedPswd.Error())
		case errors.Is(err, entities.ErrCountryLength):
			return errorResponse(c, http.StatusBadRequest, err, entities.ErrCountryLength.Error())
		case errors.Is(err, entities.ErrExistingEmail):
			return errorResponse(c, http.StatusConflict, err, entities.ErrExistingEmail.Error())
		case errors.Is(err, entities.ErrInsertUser):
			return errorResponse(c, http.StatusFailedDependency, err, entities.ErrInsertUser.Error())
		case errors.Is(err, entities.ErrPublishMsg):
			return errorResponse(c, http.StatusFailedDependency, err, entities.ErrPublishMsg.Error())
		}

		return errorResponse(c, http.StatusInternalServerError, err, "")
	}

	resp := userRespFromUser(user)
	return c.JSON(http.StatusCreated, resp)
}
