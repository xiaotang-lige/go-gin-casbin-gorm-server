package model

type UserConfig struct {
	TChatModel
	UserId   string `json:"userId" form:"userId" binding:"required"`
	UserName string `json:"userName" form:"userName" `
	Password string `json:"password" form:"password" binding:"required"`
	Token    string `json:"token" form:"token" `
	Status   int    `json:"status"  form:"status" ` //身份
	LogoPath string `json:"logoPath" form:"logoPath" `
}

func (*UserConfig) TableName() string {
	return "user_configs"
}
