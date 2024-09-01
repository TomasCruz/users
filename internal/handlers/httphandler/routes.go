package httphandler

func (h HTTPHandler) bindRoutes() {
	h.e.GET("/health", h.HealthHandler)
	// h.e.PUT("/users", h.CreateUserHandler)
	// h.e.GET("/users/:user-id", h.GetUserHandler)
	// h.e.GET("/users", h.ListUserHandler)
	// h.e.POST("/users/:user-id", h.UpdateUserHandler)
	// h.e.DELETE("/users/:user-id", h.DeleteUserHandler)
}
