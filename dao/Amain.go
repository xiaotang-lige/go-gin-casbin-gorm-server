package dao

import (
	"log"
	"messageServe/linker"
	"messageServe/model"
)

type dao struct {
	UserConfig *userConfig
	Contacts   *contacts
}

var Api = new(dao)

func Main() {
	err := linker.Api.MysqlDb.AutoMigrate(
		&model.UserConfig{},
		&model.Contacts{},
	)
	if err != nil {
		log.Println(err)
	}
}
