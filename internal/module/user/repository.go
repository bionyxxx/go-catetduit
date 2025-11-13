package user

type Repository interface {
	GetUserByID(id int) (*User, error)
}
