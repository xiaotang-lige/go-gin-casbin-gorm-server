package tool

import (
	"github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"log"
	"messageServe/linker"
	"messageServe/model"
)

func CasbinHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := Api.Token.GetClaimsModel(c).StatusId // 想要访问资源的用户。
		obj := c.Request.URL.Path
		// 将被访问的资源。
		act := c.Request.Method // 用户对资源执行的操作。
		b, err := casbinDb().Enforce(sub, obj, act)
		if err != nil {
			Api.Response.Write(c, &model.Response{State: 400, Err: "非法访问！"})
			c.Abort()
			return
		}
		if !b {
			Api.Response.Write(c, &model.Response{State: 401, Err: "无权限！"})
			c.Abort()
			return
		}
		c.Next()
	}
}
func casbinDb() *casbin.Enforcer {
	db, err := gormadapter.NewAdapterByDB(linker.Api.MysqlDb)
	m, err := casbinModel.NewModelFromString(`
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
	err = e.LoadPolicy()
	if err != nil {
		log.Println("casbin出錯:", err)
	}
	return e
}
