package dao

import (
	"log"
	"messageServe/linker"
	"messageServe/model"
	"testing"
)

func TestDao(t *testing.T) {
	linker.Main()
	data := &model.UserConfig{
		UserId:   "123",
		Password: "123",
		Token:    "1",
		Status:   1,
		LogoPath: "1",
	}
	err := Api.UserConfig.Create(data)
	//var i interface{}
	//db, i := Api.UserConfig.Query(data)
	//if err := db.Error; err != nil {
	//	log.Println(err)
	//}
	log.Println(err)
}
func TestDaoContacts(t *testing.T) {
	linker.Main()
	d := &model.Contacts{
		A: "xiaowang",
		B: "dawang",
	}
	db := Api.Contacts.Create(d)
	if db.Error != nil {
		log.Println(db.Error)
	}
}
func TestOnlineList_Get(t *testing.T) {
	linker.Main()
	Api.OnlineList.Add("asdsadasda", "asdsad", 0)

	//Api.OnlineList.Get()
}
