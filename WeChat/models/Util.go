package models

type ResponseResult struct {
	Code    int64
	Success bool
	Message string
	Data    interface{}
	Data62  string
}

type ResponseResult2 struct {
	Code     int64
	Success  bool
	Message  string
	Data     interface{}
	Data62   string
	DeviceId string
}

