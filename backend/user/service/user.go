package service

import "user/model/dto"

type UserService interface {
	Login(string, string) (dto.UserLoginResponse, error)
	Register(string, string) (dto.UserRegisterResponse, error)
	Info(string) (string, error)

	/* 以下需要鉴权 */

	ChangeAvatar(string, string) (string, error)
}
