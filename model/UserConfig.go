package model

import "gorm.io/gorm"

type UserConfig struct {
	*gorm.Model
	UserId   string `json:"userId" form:"userId" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Token    string `json:"token" form:"token" `
	Status   int32  `json:"status"  form:"status" ` //身份
	LogoPath string `json:"logoPath" form:"logoPath" `
}
