package dto

type ProgramDTO struct {
	ID          uint   `gorm:"primary_key" json:"id,omitempty"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
	Cover       string `gorm:"cover"`
	Category    string `gorm:"category"`
}
