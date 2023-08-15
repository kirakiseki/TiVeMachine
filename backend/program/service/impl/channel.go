package impl

import (
	"program/dao"
	"program/model/dto"
	"program/utils"
)

type ChannelServiceImpl struct{}

func (c *ChannelServiceImpl) List() (dto.Response, error) {
	resp := dto.Response{}

	list, err := dao.GetChannelList()
	if err != nil {
		utils.Fail(&resp, 50027, "查询失败")
		return resp, err
	}

	utils.Success(&resp)
	listResp := dto.ListResponse{
		List: list,
	}
	resp.Data = listResp
	return resp, nil
}

func (c *ChannelServiceImpl) Info(channelID uint) (dto.Response, error) {
	resp := dto.Response{}

	if !dao.HasChannelByID(channelID) {
		utils.Fail(&resp, 50028, "频道不存在")
		return resp, nil
	}

	channel, err := dao.GetChannelInfo(channelID)
	if err != nil {
		utils.Fail(&resp, 50029, "获取频道信息失败")
		return resp, err
	}

	utils.Success(&resp)
	infoResp := dto.ChannelInfoResponse{
		Info: channel,
	}
	resp.Data = infoResp
	
	return resp, nil
}
