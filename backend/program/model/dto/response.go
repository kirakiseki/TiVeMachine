package dto

type Response struct {
	StatusCode int    `json:"code"`
	Message    string `json:"msg"`
	Status     string `json:"status"`
	Data       any    `json:"data,omitempty"`
}

type ListResponse struct {
	List []uint `json:"list"`
}

type ProgramInfoResponse struct {
	Info ProgramDTO `json:"info"`
}

type ChannelInfoResponse struct {
	Info ChannelDTO `json:"info"`
}
