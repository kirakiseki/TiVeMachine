package dao

import (
	"program/model/dto"
	"program/setup"
)

type ProgramPO struct {
	dto.ProgramDTO
}

func (p ProgramPO) TableName() string {
	return "program"
}

func GetProgramList() ([]uint, error) {
	var programList []uint
	err := setup.Inst.DB.Model(&ProgramPO{}).Select("id").Find(&programList).Error
	return programList, err
}

func GetProgramListByPeriod(startTime, endTime uint) ([]uint, error) {
	var programList []uint
	err := setup.Inst.DB.Model(&ProgramPO{}).Joins("INNER JOIN schedule ON program.id = schedule.program_id").Where("start_time >= ? AND end_time <= ?", startTime, endTime).Distinct("program.id").Find(&programList).Error
	return programList, err
}

func GetProgramListByCategory(category string) ([]uint, error) {
	var programList []uint
	err := setup.Inst.DB.Model(&ProgramPO{}).Where("category LIKE ?", "%"+category+"%").Select("id").Find(&programList).Error
	return programList, err
}

func GetProgramListByChannel(channelID uint) ([]uint, error) {
	var programList []uint
	err := setup.Inst.DB.Model(&ProgramPO{}).Joins("INNER JOIN schedule ON program.id = schedule.program_id").Where("channel_id = ?", channelID).Distinct("program.id").Find(&programList).Error
	return programList, err
}

func HasProgramByID(programID uint) bool {
	var count int64
	setup.Inst.DB.Model(&ProgramPO{}).Where("id = ?", programID).Count(&count)
	return count > 0
}

func HasProgramByName(programName string) bool {
	var count int64
	setup.Inst.DB.Model(&ProgramPO{}).Where("name = ?", programName).Count(&count)
	return count > 0
}

func GetProgramInfo(programID uint) (dto.ProgramDTO, error) {
	var programInfo dto.ProgramDTO
	err := setup.Inst.DB.Model(&ProgramPO{}).Where("id = ?", programID).Find(&programInfo).Error
	return programInfo, err
}

func SearchProgram(programName string) ([]uint, error) {
	var programList []uint
	err := setup.Inst.DB.Model(&ProgramPO{}).Where("name LIKE ?", "%"+programName+"%").Select("id").Find(&programList).Error
	return programList, err
}

func AddProgram(program dto.ProgramDTO) error {
	return setup.Inst.DB.Model(&ProgramPO{}).Create(&ProgramPO{
		ProgramDTO: program,
	}).Error
}
