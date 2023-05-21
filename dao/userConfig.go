package dao

import (
	"gorm.io/gorm"
	"messageServe/linker"
	"messageServe/model"
)

type userConfig struct {
}

func (*userConfig) Create(data *model.UserConfig) *gorm.DB {
	return linker.Db.MysqlDb.Create(&data)
}
func (*userConfig) Query(data *model.UserConfig) {
	linker.Db.MysqlDb.Find(&data)
}
