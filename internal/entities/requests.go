package entities

type CreateUserReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PswdHash  string `json:"pswd_hash"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

type UpdateUserReq struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	PswdHash  *string `json:"pswd_hash,omitempty"`
	Email     *string `json:"email,omitempty"`
	Country   *string `json:"country,omitempty"`
}
