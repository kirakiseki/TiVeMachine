package service

import "subscribe/model/dto"

type SubscribeService interface {
	/* 以下需要鉴权 */

	SubscriptionList(uint) (dto.Response, error)
	Subscribe(uint, uint, uint) (dto.Response, error)
	Unsubscribe(uint, uint) (dto.Response, error)
	SubscriptionInfo(uint) (dto.Response, error)
}
