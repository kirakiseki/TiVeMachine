package dto

type ScheduleDTO struct {
	ID        uint `gorm:"primary_key"`
	ProgramID uint `gorm:"program_id"`
	ChannelID uint `gorm:"channel_id"`
	StartTime uint `gorm:"start_time"`
	EndTime   uint `gorm:"end_time"`
}
