package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
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

func SetAvatar(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "SetAvatar"})
}

type SetDescriptionRequest struct {
	Description string `json:"description"`
}

func SetDescription(c *gin.Context) {
	var req SetDescriptionRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 50019, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := userServiceImpl.SetDescription(c.GetUint("userID"), req.Description)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("SetDescription")
	}

	c.JSON(http.StatusOK, resp)
}

type SetSexRequest struct {
	Sex uint `json:"sex"`
}

func SetSex(c *gin.Context) {
	var req SetSexRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 50020, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := userServiceImpl.SetSex(c.GetUint("userID"), req.Sex)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("SetSex")
	}

	c.JSON(http.StatusOK, resp)
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

func UploadAvatar(c *gin.Context) {
	data, err := getFile(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := userServiceImpl.Upload(data, "avatar")
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("UploadAvatar")
	}
	c.JSON(http.StatusOK, resp)
}

func UploadCover(c *gin.Context) {
	data, err := getFile(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := userServiceImpl.Upload(data, "cover")
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("UploadCover")
	}
	c.JSON(http.StatusOK, resp)
}

func getFile(c *gin.Context) (*multipart.FileHeader, error) {
	data, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}
	return data, nil
}
