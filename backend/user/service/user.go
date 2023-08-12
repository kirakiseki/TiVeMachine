package service

type UserService interface {
	Login(string, string) (string, error)
	Register(string, string) (string, error)
	Info(string) (string, error)

	/* 以下需要鉴权 */

	ChangeAvatar(string, string) (string, error)
}
