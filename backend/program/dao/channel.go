package dao

import (
	"program/model/dto"
	"program/setup"
)

type ChannelPO struct {
	dto.ChannelDTO
}

func (c ChannelPO) TableName() string {
	return "channel"
}

func GetChannelList() ([]uint, error) {
	var channelList []uint
	err := setup.Inst.DB.Model(&ChannelPO{}).Select("id").Find(&channelList).Error
	return channelList, err
}

func HasChannelByID(channelID uint) bool {
	var count int64
	setup.Inst.DB.Model(&ChannelPO{}).Where("id = ?", channelID).Count(&count)
	return count > 0
}

func GetChannelInfo(channelID uint) (dto.ChannelDTO, error) {
	var channelInfo dto.ChannelDTO
	err := setup.Inst.DB.Model(&ChannelPO{}).Where("id = ?", channelID).Find(&channelInfo).Error
	return channelInfo, err
}
