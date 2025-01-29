package service

type AuthUseCase interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
}
