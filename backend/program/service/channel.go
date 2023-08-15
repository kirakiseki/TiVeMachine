package service

import "program/model/dto"

type ChannelService interface {
	List() (dto.Response, error)
	Info(uint) (dto.Response, error)
}
