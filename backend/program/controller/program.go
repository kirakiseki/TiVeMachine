package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"program/model/dto"
	"program/service/impl"
	"program/setup"
)

var programServiceImpl = impl.ProgramServiceImpl{}
var channelServiceImpl = impl.ChannelServiceImpl{}

func List(c *gin.Context) {
	resp, err := programServiceImpl.List()
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("List")
	}

	c.JSON(http.StatusOK, resp)
}

func ChannelList(c *gin.Context) {
	resp, err := channelServiceImpl.List()
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("ChannelList")
	}

	c.JSON(http.StatusOK, resp)
}

type InfoRequest struct {
	ProgramID uint `json:"program_id"`
}

func Info(c *gin.Context) {
	var req InfoRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Info")
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}
	if req.ProgramID == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": "program_id 不能为空", "status": "fail"})
		return
	}

	resp, err := programServiceImpl.Info(req.ProgramID)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Info")
	}

	c.JSON(http.StatusOK, resp)
}

type ChannelInfoRequest struct {
	ChannelID uint `json:"channel_id"`
}

func ChannelInfo(c *gin.Context) {
	var req ChannelInfoRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("ChannelInfo")
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}
	if req.ChannelID == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": "channel_id 不能为空", "status": "fail"})
		return
	}

	resp, err := channelServiceImpl.Info(req.ChannelID)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("ChannelInfo")
	}

	c.JSON(http.StatusOK, resp)
}

type ListByPeriodRequest struct {
	StartTime uint `json:"start_time"`
	EndTime   uint `json:"end_time"`
}

func ListByPeriod(c *gin.Context) {
	var req ListByPeriodRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("ListByPeriod")
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := programServiceImpl.ListByPeriod(req.StartTime, req.EndTime)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("ListByPeriod")
	}

	c.JSON(http.StatusOK, resp)
}

type ListByCategoryRequest struct {
	Category string `json:"category"`
}

func ListByCategory(c *gin.Context) {
	var req ListByCategoryRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("ListByCategory")
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := programServiceImpl.ListByCategory(req.Category)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("ListByCategory")
	}

	c.JSON(http.StatusOK, resp)
}

type ListByChannelRequest struct {
	ChannelID uint `json:"channel_id"`
}

func ListByChannel(c *gin.Context) {
	var req ListByChannelRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("ListByChannel")
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := programServiceImpl.ListByChannel(req.ChannelID)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("ListByChannel")
	}

	c.JSON(http.StatusOK, resp)
}

type SearchRequest struct {
	Keyword string `json:"keyword"`
}

func Search(c *gin.Context) {
	var req SearchRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Search")
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := programServiceImpl.Search(req.Keyword)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Search")
	}

	c.JSON(http.StatusOK, resp)
}

type AddRequest struct {
	dto.ProgramDTO
}

func Add(c *gin.Context) {
	var req AddRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Add")
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := programServiceImpl.Add(req.ProgramDTO)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Add")
	}

	c.JSON(http.StatusOK, resp)
}
