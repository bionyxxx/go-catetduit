package user

type Repository interface {
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) (*User, error)
}
