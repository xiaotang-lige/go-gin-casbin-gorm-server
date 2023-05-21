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
		Password: "1",
		Token:    "1",
		Status:   1,
		LogoPath: "1",
	}
	err := DaoApi.Create(data).Error
	if err != nil {
		log.Println(err)
	}
}
