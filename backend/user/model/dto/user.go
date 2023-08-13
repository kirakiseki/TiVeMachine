package dto

type UserLoginDTO struct {
	ID       uint
	Password string
}

type UserInfoDTO struct {
	ID          uint   `json:"user_id"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar,omitempty"`
	Description string `json:"description,omitempty"`
	Sex         uint   `json:"sex,omitempty"`
}

type UserDTO struct {
	ID          uint `gorm:"primary_key"`
	Username    string
	Avatar      string
	Description string
	Sex         uint
	Password    string
}
