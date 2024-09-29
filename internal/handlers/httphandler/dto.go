package httphandler

import "github.com/TomasCruz/users/internal/core/entities"

func userDTOFromCreateUserReq(req CreateUserReq) entities.UserDTO {
	return entities.UserDTO{
		FirstName: &req.FirstName,
		LastName:  &req.LastName,
		PswdHash:  &req.PswdHash,
		Email:     &req.Email,
		Country:   &req.Country,
	}
}
