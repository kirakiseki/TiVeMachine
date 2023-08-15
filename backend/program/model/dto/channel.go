package dto

type ChannelDTO struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
	Cover       string `gorm:"cover"`
}
