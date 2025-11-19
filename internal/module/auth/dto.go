package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name                 string `json:"name" validate:"required,min=2,max=100"`
	Phone                string `json:"phone" validate:"required,phone,unique=users.phone"`
	Email                string `json:"email" validate:"required,email,unique=users.email"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
}
