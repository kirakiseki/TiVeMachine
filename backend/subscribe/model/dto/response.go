package dto

type Response struct {
	StatusCode int    `json:"code"`
	Message    string `json:"msg"`
	Status     string `json:"status"`
	Data       any    `json:"data,omitempty"`
}

type SubscriptionListResponse struct {
	SubscriptionList []uint `json:"subscription_list"`
}
