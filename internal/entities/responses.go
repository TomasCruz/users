package entities

import (
	"time"

	"github.com/google/uuid"
)

type ErrResp struct {
	Msg string `json:"errorMessage" example:"A horrible, terrible, absolutely awful error"`
}

type UserResp struct {
	UserID    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	PswdHash  string    `json:"pswd_hash"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// conversion warning (S1016) should go away after introducing i.e. soft delete or anything else "invisible" to API user
func UserRespFromUser(user User) UserResp {
	return UserResp{
		UserID:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		PswdHash:  user.PswdHash,
		Email:     user.Email,
		Country:   user.Country,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
