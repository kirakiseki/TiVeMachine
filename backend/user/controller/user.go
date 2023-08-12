package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user/service/impl"
)

var userServiceImpl = impl.UserServiceImpl{}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login"})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register"})
}

func Info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Info"})
}
