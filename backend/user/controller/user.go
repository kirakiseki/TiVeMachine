package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"user/service/impl"
	"user/setup"
)

var userServiceImpl = impl.UserServiceImpl{}

func Login(c *gin.Context) {
	username, password, err := getUsernameAndPassword(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := userServiceImpl.Login(username, password)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Login")
	}

	c.JSON(http.StatusOK, resp)
}

func Register(c *gin.Context) {
	username, password, err := getUsernameAndPassword(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := userServiceImpl.Register(username, password)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Register")
	}

	c.JSON(http.StatusOK, resp)
}

func Info(c *gin.Context) {
	resp, err := userServiceImpl.Info(c.GetUint("userID"))
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Info")
	}

	c.JSON(http.StatusOK, resp)
}

func ChangeAvatar(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ChangeAvatar"})
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getUsernameAndPassword(c *gin.Context) (string, string, error) {
	json := UserRequest{}

	err := c.ShouldBindJSON(&json)
	if err != nil {
		return "", "", err
	}

	username, password := json.Username, json.Password
	if username == "" || password == "" {
		return "", "", errors.New("用户名或密码不能为空")
	}

	return username, password, nil
}
