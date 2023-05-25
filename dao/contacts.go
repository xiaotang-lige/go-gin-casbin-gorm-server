package dao

import (
	"gorm.io/gorm"
	"messageServe/linker"
	"messageServe/model"
)

type contacts struct {
}

func (*contacts) TableName() string {
	m := model.Contacts{}
	return m.TableName()
}

// QueryAll 选择查询所有
func (t *contacts) QueryAll(config *model.UserConfig, state int) (listData []*model.Contacts, db *gorm.DB) {
	db = linker.Api.MysqlDb.Table(t.TableName()).Select("(a = ? or b = ?) and state = ?", config.UserId, config.UserId, state).Find(&listData)
	return listData, db
}
func (t *contacts) Query(config *model.Contacts) (data *model.Contacts, db *gorm.DB) {
	db = linker.Api.MysqlDb.Table(t.TableName()).Select("a = ? or b = ?").First(&data)
	return data, db
}
func (t *contacts) QueryState(config *model.Contacts) (data *model.Contacts, db *gorm.DB) {
	db = linker.Api.MysqlDb.Table(t.TableName()).Select("a = ? or b = ?").First(&data)
	return data, db
}
func (t *contacts) Create(config *model.Contacts) *gorm.DB {
	config.State = 1
	return linker.Api.MysqlDb.Table(t.TableName()).Create(&config)
}
