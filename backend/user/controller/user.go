package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user/service/impl"
	"user/setup"
)

var userServiceImpl = impl.UserServiceImpl{}

func Login(c *gin.Context) {
	username, password, err := getUsernameAndPassword(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 50009, "message": err.Error()})
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
		c.JSON(http.StatusOK, gin.H{"status_code": 50009, "message": err.Error()})
		return
	}

	resp, err := userServiceImpl.Register(username, password)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Register")
	}

	c.JSON(http.StatusOK, resp)
}

func Info(c *gin.Context) {
	usedID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Info")
	}

	resp, err := userServiceImpl.Info(uint(usedID))
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Info")
	}

	c.JSON(http.StatusOK, resp)
}

func ChangeAvatar(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ChangeAvatar"})
}

func getUsernameAndPassword(c *gin.Context) (string, string, error) {
	username := c.Query("username")
	password := c.Query("password")

	if username == "" || password == "" {
		return "", "", errors.New("用户名或密码不能为空")
	}

	return username, password, nil
}
