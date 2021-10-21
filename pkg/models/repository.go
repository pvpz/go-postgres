package models

type Repository interface {
	InsertUser(User) (int, error)
	GetUser(int) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(int, User) (int64, error)
	DeleteUser(int) (int64, error)
}
