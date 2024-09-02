package httphandler

import (
	_ "github.com/TomasCruz/users/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (h HTTPHandler) bindRoutes() {
	h.e.GET("/swagger/*any", echoSwagger.WrapHandler)
	h.e.GET("/health", h.HealthHandler)
	h.e.GET("/users/:user-id", h.GetUserHandler)
	// h.e.GET("/users", h.ListUserHandler)
	// h.e.PUT("/users", h.CreateUserHandler)
	// h.e.POST("/users/:user-id", h.UpdateUserHandler)
	// h.e.DELETE("/users/:user-id", h.DeleteUserHandler)
}
