package models

type BaseResponse struct {
	ResultMessage string      `json:"resultMessage"`
	Data          interface{} `json:"data,omitempty"`
}

type BaseResponseWithErrors struct {
	ResultMessage string `json:"resultMessage"`
}
