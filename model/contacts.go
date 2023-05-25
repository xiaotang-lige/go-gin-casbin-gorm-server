package model

type Contacts struct {
	TChatModel
	A     string
	B     string
	State int //1:好友状态 2：删除状态
	T1    string
	T2    string
	T3    string
}

func (*Contacts) TableName() string {
	return "contacts"
}
