package types

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type RegisterUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,len=8"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
