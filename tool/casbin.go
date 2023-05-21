package tool

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"log"
	"messageServe/linker"
)

func CasbinHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := "alice" // 想要访问资源的用户。
		obj := c.Request.URL.Path
		// 将被访问的资源。
		act := c.Request.Method // 用户对资源执行的操作。
		log.Println(sub, obj, act)
		b, err := casbinDb().Enforce(sub, obj, act)
		if err != nil {
			log.Println(err)
		}
		if !b {
			c.JSON(401, gin.H{
				"message": "error:無權限！",
				"data":    "",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
func casbinDb() *casbin.Enforcer {
	db, err := gormadapter.NewAdapterByDB(linker.Db.MysqlDb)
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
