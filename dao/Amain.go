package dao

import (
	"log"
	"messageServe/linker"
	"messageServe/model"
)

type dao struct {
	userConfig
}

var DaoApi = new(dao)

func Main() {
	autoCreateTable(
		&model.UserConfig{},
	)
}

// 創建表
func autoCreateTable(i interface{}) {
	err := linker.Db.MysqlDb.AutoMigrate(i)
	if err != nil {
		log.Println(err)
	}
}
