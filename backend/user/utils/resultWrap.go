package utils

import "user/model/dto"

func Success(resp *dto.Response) {
	resp.StatusCode = 200
	resp.Message = "Success"
}

func Fail(resp *dto.Response, code int, message string) {
	resp.StatusCode = code
	resp.Message = message
}
