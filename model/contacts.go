package model

type Contacts struct {
	TChatModel
	A         string `json:"a" gorm:"column:a"`
	ALogoPath string `json:"aLogoPath" gorm:"column:aLogoPath"`
	B         string `json:"b" gorm:"column:b"`
	BLogoPath string `json:"bLogoPath" gorm:"column:bLogoPath"`
	State     int    `json:"stateId" gorm:"column:state"` //1:好友状态 2：删除状态
	T1        string `json:"t1" gorm:"column:t1"`
	T2        string `json:"t2" gorm:"column:t2"`
	T3        string `json:"t3" gorm:"column:t3"`
}

func (*Contacts) TableName() string {
	return "contacts"
}
