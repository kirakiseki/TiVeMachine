package service

import (
	"mime/multipart"
	"user/model/dto"
)

type UserService interface {
	Login(string, string) (dto.Response, error)
	Register(string, string) (dto.Response, error)
	Upload(*multipart.FileHeader, string) (dto.Response, error)

	/* 以下需要鉴权 */

	Info() (dto.Response, error)
	ChangeAvatar(string, string) (dto.Response, error)
}
