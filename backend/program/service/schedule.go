package service

import "program/model/dto"

type ScheduleService interface {
	ScheduleInfo(uint) (dto.Response, error)
	ScheduleList() (dto.Response, error)

	/* 以下需要鉴权 */
}
