package dto

type Response struct {
	StatusCode int    `json:"code"`
	Message    string `json:"msg"`
	Status     string `json:"status"`
	Data       any    `json:"data,omitempty"`
}

type UserRegisterResponse struct {
	UserID uint `json:"user_id"`
}

type UserLoginResponse struct {
	Token  string `json:"token"`
	UserID uint   `json:"user_id"`
}

type UserInfoResponse UserDTO
