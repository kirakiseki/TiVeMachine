package impl

import (
	"errors"
	"subscribe/dao"
	"subscribe/model/dto"
	"subscribe/utils"
)

type SubscribeServiceImpl struct{}

func (s *SubscribeServiceImpl) SubscriptionList(userID uint) (dto.Response, error) {
	resp := dto.Response{}

	if !dao.HasSubscription(userID) {
		utils.Success(&resp)
		loginResp := dto.SubscriptionListResponse{
			SubscriptionList: make([]uint, 0),
		}
		resp.Data = loginResp

		return resp, nil
	}

	list, err := dao.GetSubscriptionListByUserID(userID)
	if err != nil {
		utils.Fail(&resp, 50011, "查询失败")
		return resp, err
	}

	utils.Success(&resp)
	var scheduleID []uint
	for _, v := range list {
		scheduleID = append(scheduleID, v.ScheduleID)
	}

	loginResp := dto.SubscriptionListResponse{
		SubscriptionList: scheduleID,
	}
	resp.Data = loginResp

	return resp, nil
}

func (s *SubscribeServiceImpl) Subscribe(userID, scheduleID uint) (dto.Response, error) {
	resp := dto.Response{}

	if dao.HasSubscriptionForScheduleID(userID, scheduleID) {
		utils.Fail(&resp, 50015, "已订阅该节目")
		return resp, errors.New("已订阅该节目")
	}

	err := dao.Subscribe(userID, scheduleID)
	if err != nil {
		utils.Fail(&resp, 50012, "订阅失败")
		return resp, err
	}

	utils.Success(&resp)

	return resp, nil
}

func (s *SubscribeServiceImpl) Unsubscribe(userID, scheduleID uint) (dto.Response, error) {
	resp := dto.Response{}

	if !dao.HasSubscriptionForScheduleID(userID, scheduleID) {
		utils.Fail(&resp, 50013, "没有订阅该节目")
		return resp, errors.New("没有订阅该节目")
	}

	err := dao.Unsubscribe(userID, scheduleID)
	if err != nil {
		utils.Fail(&resp, 50014, "取消订阅失败")
		return resp, err
	}

	utils.Success(&resp)

	return resp, nil
}
