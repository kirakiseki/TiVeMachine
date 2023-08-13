package impl

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"user/dao"
	"user/model/dto"
	"user/routers/middleware"
	"user/setup"
	"user/utils"
)

type UserServiceImpl struct{}

func (u *UserServiceImpl) Login(username string, password string) (dto.UserLoginResponse, error) {
	resp := dto.UserLoginResponse{}

	user, err := dao.GetUserLoginDTOByUsername(username)
	if err != nil {
		utils.Fail(&resp.Response, 50002, "用户不存在")
		return resp, err
	}

	encryptPassword, err := utils.Encrypt(password)
	if err != nil {
		utils.Fail(&resp.Response, 50004, err.Error())
		return resp, err
	}

	if user.Password != encryptPassword {
		utils.Fail(&resp.Response, 50003, "密码错误")
		return resp, errors.New("密码错误")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, middleware.JWTClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	tokenSigned, err := token.SignedString([]byte(setup.Config.JWTSecret))
	if err != nil {
		utils.Fail(&resp.Response, 50005, "生成Token失败")
		return resp, err
	}

	utils.Success(&resp.Response)
	resp.UserID = user.ID
	resp.Token = tokenSigned

	return resp, nil
}

func (u *UserServiceImpl) Register(username string, password string) (dto.UserRegisterResponse, error) {
	resp := dto.UserRegisterResponse{}

	if dao.HasUserByUsername(username) {
		utils.Fail(&resp.Response, 50006, "用户名已存在")
		return resp, errors.New("用户名已存在")
	}

	encryptPassword, err := utils.Encrypt(password)
	if err != nil {
		utils.Fail(&resp.Response, 50004, err.Error())
		return resp, err
	}

	user := dto.UserDTO{
		Username: username,
		Password: encryptPassword,
	}

	err = dao.RegisterUser(user)
	if err != nil {
		utils.Fail(&resp.Response, 50007, "注册失败")
		return resp, err
	}

	utils.Success(&resp.Response)
	resp.UserID = user.ID

	return resp, nil
}

func (u *UserServiceImpl) Info(id uint) (dto.UserInfoResponse, error) {
	resp := dto.UserInfoResponse{}

	if !dao.HasUserByID(id) {
		utils.Fail(&resp.Response, 50002, "用户不存在")
		return resp, errors.New("用户不存在")
	}

	user, err := dao.GetUserInfoDTOByID(id)
	if err != nil {
		utils.Fail(&resp.Response, 50008, "获取用户信息失败")
		return resp, err
	}

	utils.Success(&resp.Response)
	resp.UserInfoDTO = user

	return resp, nil
}

func (u *UserServiceImpl) ChangeAvatar(username string, avatar string) (string, error) {
	return "ChangeAvatar", nil
}
