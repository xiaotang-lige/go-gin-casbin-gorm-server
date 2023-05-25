package model

type Response struct {
	Len   int         `json:"len" form:"len"`
	State int         `json:"state" form:"state"`
	Err   string      `json:"err" form:"err"`
	Data  interface{} `json:"data" form:"data"`
}
