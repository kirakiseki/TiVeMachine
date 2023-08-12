package service

type UserService interface {
	// TODO FIX TYPE
	Login(string, string) (string, error)
	Register(string, string) (string, error)
	Info(string) (string, error)
}
