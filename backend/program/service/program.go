package service

import "program/model/dto"

type ProgramService interface {
	List() (dto.Response, error)
	ListByPeriod(uint, uint) (dto.Response, error)
	ListByCategory(string) (dto.Response, error)
	ListByChannel(uint) (dto.Response, error)
	Info(uint) (dto.Response, error)
	Search(string) (dto.Response, error)
	Add(dto.ProgramDTO) (dto.Response, error)

	/* 以下需要鉴权 */
}
