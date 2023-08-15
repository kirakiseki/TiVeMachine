package dao

import (
	"user/model/dto"
	"user/setup"
)

type UserPO struct {
	dto.UserDTO
}

func (u UserPO) TableName() string {
	return "user"
}

func GetUserLoginDTOByUsername(username string) (dto.UserLoginDTO, error) {
	user := dto.UserDTO{}

	err := setup.Inst.DB.Model(&UserPO{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return dto.UserLoginDTO{}, err
	}

	userLogin := dto.UserLoginDTO{
		ID:       user.ID,
		Password: user.Password,
	}

	return userLogin, nil
}

func GetUserInfoDTOByID(id uint) (dto.UserInfoDTO, error) {
	user := dto.UserDTO{}

	err := setup.Inst.DB.Model(&UserPO{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return dto.UserInfoDTO{}, err
	}

	userInfo := dto.UserInfoDTO{
		ID:          user.ID,
		Username:    user.Username,
		Avatar:      user.Avatar,
		Description: user.Description,
		Sex:         user.Sex,
	}

	return userInfo, nil
}

func HasUserByUsername(username string) bool {
	var count int64

	setup.Inst.DB.Model(&UserPO{}).Where("username = ?", username).Count(&count)

	return count > 0
}

func HasUserByID(id uint) bool {
	var count int64

	setup.Inst.DB.Model(&UserPO{}).Where("id = ?", id).Count(&count)

	return count > 0
}

func RegisterUser(user dto.UserDTO) error {
	return setup.Inst.DB.Create(&UserPO{
		UserDTO: user,
	}).Error
}

func UpdateUserAvatar(id uint, avatar string) error {
	return setup.Inst.DB.Model(&UserPO{}).Where("id = ?", id).Update("avatar", avatar).Error
}

func UpdateUserDescription(id uint, description string) error {
	return setup.Inst.DB.Model(&UserPO{}).Where("id = ?", id).Update("description", description).Error
}

func UpdateUserSex(id uint, sex uint) error {
	return setup.Inst.DB.Model(&UserPO{}).Where("id = ?", id).Update("sex", sex).Error
}
