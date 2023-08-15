package dto

type SubscribeDTO struct {
	ID         uint `gorm:"primary_key"`
	UserID     uint `gorm:"user_id"`
	ScheduleID uint `gorm:"schedule_id"`
}
