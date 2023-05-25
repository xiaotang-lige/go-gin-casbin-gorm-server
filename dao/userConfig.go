package dao

import (
	"gorm.io/gorm"
	"messageServe/linker"
	"messageServe/model"
)

type userConfig struct {
}

// Create 成功：1,并err nil
func (*userConfig) Create(data *model.UserConfig) *gorm.DB {
	if linker.Api.MysqlDb.Where("user_id = ?", data.UserId).First(&data).RowsAffected == 1 {
		return nil
	}
	return linker.Api.MysqlDb.Create(&data)
}
func (*userConfig) Query(data *model.UserConfig) (db *gorm.DB, d *model.UserConfig) {
	db = linker.Api.MysqlDb.Where("user_id = ? AND password = ?", data.UserId, data.Password).First(&d)
	return db, d
}
