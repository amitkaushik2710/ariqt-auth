package models

type TestReq struct {
	Data bool `json:"data" binding:"required"`
}
