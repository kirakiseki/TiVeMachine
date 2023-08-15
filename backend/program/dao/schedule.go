package dao

import (
	"program/model/dto"
	"program/setup"
)

type SchedulePO struct {
	dto.ScheduleDTO
}

func (s SchedulePO) TableName() string {
	return "schedule"
}

func GetScheduleInfo(scheduleID uint) (dto.ScheduleDTO, error) {
	var schedule dto.ScheduleDTO
	err := setup.Inst.DB.Model(&SchedulePO{}).Where("id = ?", scheduleID).First(&schedule).Error
	return schedule, err
}

func GetScheduleList() ([]uint, error) {
	var scheduleList []uint
	err := setup.Inst.DB.Model(&SchedulePO{}).Select("id").Find(&scheduleList).Error
	return scheduleList, err
}
