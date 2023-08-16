package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"subscribe/service/impl"
	"subscribe/setup"
)

var subscribeServiceImpl = impl.SubscribeServiceImpl{}

func SubscriptionList(c *gin.Context) {
	resp, err := subscribeServiceImpl.SubscriptionList(c.GetUint("userID"))
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("SubscriptionList")
	}

	c.JSON(http.StatusOK, resp)
}

func Subscribe(c *gin.Context) {
	scheduleID, alarmTime, err := getScheduleID(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := subscribeServiceImpl.Subscribe(c.GetUint("userID"), scheduleID, alarmTime)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Subscribe")
	}

	c.JSON(http.StatusOK, resp)
}

func Unsubscribe(c *gin.Context) {
	scheduleID, _, err := getScheduleID(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 50009, "msg": err.Error(), "status": "fail"})
		return
	}

	resp, err := subscribeServiceImpl.Unsubscribe(c.GetUint("userID"), scheduleID)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("Unsubscribe")
	}

	c.JSON(http.StatusOK, resp)
}

type SubscribeRequest struct {
	ScheduleID uint `json:"schedule_id"`
	AlarmTime  uint `json:"alarm_time,omitempty"`
}

func getScheduleID(c *gin.Context) (uint, uint, error) {
	var req SubscribeRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		return 0, 0, err
	}

	if req.ScheduleID == 0 {
		return 0, 0, errors.New("schedule_id is required")
	}

	return req.ScheduleID, req.AlarmTime, nil
}
