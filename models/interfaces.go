package models

// UserRepositoryInterface interface for user repository
type UserRepositoryInterface interface {
	CreateUser(u User) (User, error)
	GetUserByEmail(email string) (User, error)
}
