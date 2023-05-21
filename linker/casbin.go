package linker

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

func casbinDb() *casbin.Enforcer {
	db, err := gormadapter.NewAdapterByDB(Db.MysqlDb)
	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	e, err := casbin.NewEnforcer(m, db)
	if err != nil {
		log.Println("casbin出錯:", err)
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Println("casbin出錯:", err)
	}
	return e
}
