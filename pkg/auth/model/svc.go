package model

type AuthService interface {
	Login(email, password string) (string, error)
}
