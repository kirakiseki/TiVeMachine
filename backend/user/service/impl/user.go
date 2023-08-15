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

func (u *UserServiceImpl) Login(username string, password string) (dto.Response, error) {
	resp := dto.Response{}

	user, err := dao.GetUserLoginDTOByUsername(username)
	if err != nil {
		utils.Fail(&resp, 50002, "用户不存在")
		return resp, err
	}

	encryptPassword, err := utils.Encrypt(password)
	if err != nil {
		utils.Fail(&resp, 50004, err.Error())
		return resp, err
	}

	if user.Password != encryptPassword {
		utils.Fail(&resp, 50003, "密码错误")
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
		utils.Fail(&resp, 50005, "生成Token失败")
		return resp, err
	}

	utils.Success(&resp)
	loginResp := dto.UserLoginResponse{
		Token:  tokenSigned,
		UserID: user.ID,
	}
	resp.Data = loginResp

	return resp, nil
}

func (u *UserServiceImpl) Register(username string, password string) (dto.Response, error) {
	resp := dto.Response{}

	if dao.HasUserByUsername(username) {
		utils.Fail(&resp, 50006, "用户名已存在")
		return resp, errors.New("用户名已存在")
	}

	encryptPassword, err := utils.Encrypt(password)
	if err != nil {
		utils.Fail(&resp, 50004, err.Error())
		return resp, err
	}

	user := dto.UserDTO{
		Username: username,
		Password: encryptPassword,
	}

	err = dao.RegisterUser(user)
	if err != nil {
		utils.Fail(&resp, 50007, "注册失败")
		return resp, err
	}

	utils.Success(&resp)
	registerResp := dto.UserRegisterResponse{
		UserID: user.ID,
	}
	resp.Data = registerResp

	return resp, nil
}

func (u *UserServiceImpl) Info(id uint) (dto.Response, error) {
	resp := dto.Response{}

	if !dao.HasUserByID(id) {
		utils.Fail(&resp, 50002, "用户不存在")
		return resp, errors.New("用户不存在")
	}

	user, err := dao.GetUserInfoDTOByID(id)
	if err != nil {
		utils.Fail(&resp, 50008, "获取用户信息失败")
		return resp, err
	}

	utils.Success(&resp)
	infoResp := dto.UserInfoResponse{
		ID:          user.ID,
		Username:    user.Username,
		Avatar:      user.Avatar,
		Description: user.Description,
		Sex:         user.Sex,
	}
	resp.Data = infoResp

	return resp, nil
}

func (u *UserServiceImpl) SetAvatar(id uint, avatar string) (dto.Response, error) {
	resp := dto.Response{}

	err := dao.UpdateUserAvatar(id, avatar)
	if err != nil {
		utils.Fail(&resp, 50018, "更新头像失败")
		return resp, err
	}

	utils.Success(&resp)
	return resp, nil
}

func (u *UserServiceImpl) SetDescription(id uint, description string) (dto.Response, error) {
	resp := dto.Response{}

	err := dao.UpdateUserDescription(id, description)
	if err != nil {
		utils.Fail(&resp, 50017, "更新简介失败")
		return resp, err
	}

	utils.Success(&resp)
	return resp, nil
}

func (u *UserServiceImpl) SetSex(id, sex uint) (dto.Response, error) {
	resp := dto.Response{}

	err := dao.UpdateUserSex(id, sex)
	if err != nil {
		utils.Fail(&resp, 50016, "更新性别失败")
		return resp, err
	}

	utils.Success(&resp)
	return resp, nil
}
