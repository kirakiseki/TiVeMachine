package impl

import (
	"program/dao"
	"program/model/dto"
	"program/utils"
)

type ScheduleServiceImpl struct{}

func (s *ScheduleServiceImpl) ScheduleInfo(scheduleID uint) (dto.Response, error) {
	resp := dto.Response{}

	info, err := dao.GetScheduleInfo(scheduleID)
	if err != nil {
		utils.Fail(&resp, 50031, "查询失败")
		return resp, err
	}

	utils.Success(&resp)
	infoResp := dto.ScheduleInfoResponse{
		Info: info,
	}
	resp.Data = infoResp
	return resp, nil
}

func (s *ScheduleServiceImpl) ScheduleList() (dto.Response, error) {
	resp := dto.Response{}

	list, err := dao.GetScheduleList()
	if err != nil {
		utils.Fail(&resp, 50032, "查询失败")
		return resp, err
	}

	utils.Success(&resp)
	listResp := dto.ListResponse{
		List: list,
	}
	resp.Data = listResp
	return resp, nil
}
