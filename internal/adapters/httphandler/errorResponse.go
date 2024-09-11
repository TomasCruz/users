package httphandler

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func errorResponse(c echo.Context, errCode int, err error, msg string) error {
	if err != nil {
		if msg == "" {
			respErr := errors.Unwrap(err)
			if respErr == nil {
				respErr = err
			}

			msg = respErr.Error()
		}

		return c.JSON(errCode, ErrResp{Msg: msg})
	}

	if msg != "" {
		return c.JSON(errCode, ErrResp{Msg: msg})
	}

	return c.NoContent(errCode)
}
