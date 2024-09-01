package entities

type ErrResp struct {
	Msg string `json:"errorMessage" example:"A horrible, terrible, absolutely awful error"`
}

// type UserResp struct {
// 	UserID    uuid.UUID `json:"user_id"`
// 	FirstName string    `json:"first_name"`
// 	LastName  string    `json:"last_name"`
// 	PswdHash  string    `json:"pswd_hash"`
// 	Email     string    `json:"email"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }
