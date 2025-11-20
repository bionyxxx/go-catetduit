package user

type Repository interface {
	GetUserByID(id uint) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) (*User, error)
}
