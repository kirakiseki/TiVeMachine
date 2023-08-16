package dao

import (
	"subscribe/model/dto"
	"subscribe/setup"
)

type SubscribePO struct {
	dto.SubscribeDTO
}

func (s SubscribePO) TableName() string {
	return "subscription"
}

func GetSubscriptionListByUserID(userID uint) ([]dto.SubscribeDTO, error) {
	var subscriptions []dto.SubscribeDTO
	err := setup.Inst.DB.Model(&SubscribePO{}).Where("user_id = ?", userID).Find(&subscriptions).Error
	return subscriptions, err
}

func HasSubscription(userID uint) bool {
	var count int64
	setup.Inst.DB.Model(&SubscribePO{}).Where("user_id = ?", userID).Count(&count)
	return count > 0
}

func HasSubscriptionForScheduleID(userID, scheduleID uint) bool {
	var count int64
	setup.Inst.DB.Model(&SubscribePO{}).Where("user_id = ? AND schedule_id = ?", userID, scheduleID).Count(&count)
	return count > 0
}

func Subscribe(userID, scheduleID, alertTime uint) error {
	return setup.Inst.DB.Model(&SubscribePO{}).Create(&SubscribePO{
		SubscribeDTO: dto.SubscribeDTO{
			UserID:     userID,
			ScheduleID: scheduleID,
			AlertTime:  alertTime,
		},
	}).Error
}

func Unsubscribe(userID, scheduleID uint) error {
	return setup.Inst.DB.Model(&SubscribePO{}).Where("user_id = ? AND schedule_id = ?", userID, scheduleID).Delete(&SubscribePO{}).Error
}

func GetSubscriptionInfo(subscribeID uint) (dto.SubscribeDTO, error) {
	var subscription dto.SubscribeDTO
	err := setup.Inst.DB.Model(&SubscribePO{}).Where("id = ?", subscribeID).Find(&subscription).Error
	return subscription, err
}
