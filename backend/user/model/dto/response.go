package dto

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type UserRegisterResponse struct {
	Response
	UserID uint `json:"user_id"`
}

type UserLoginResponse struct {
	UserRegisterResponse
	Token string `json:"token"`
}

type UserInfoResponse struct {
	Response
	UserInfoDTO
}
