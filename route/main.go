package route

import (
	"github.com/gin-gonic/gin"
	"log"
)

type route struct {
	//controlApi *gin.RouterGroup
	//openApi    *gin.RouterGroup
	r *gin.Engine
}

var RouteApi = new(route)

func Main() {
	RouteApi.link()
	//無需權限
	{
		RouteApi.userConfig(RouteApi.open())
	}
	//需要權限
	{

	}
	err := RouteApi.r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
func (t *route) link() {
	t.r = gin.Default()
}

// casbin權限管理
func (t *route) control() *gin.RouterGroup {
	r := t.r.Group("")
	r.Use()
	return r
}

// 無需權限
func (t *route) open() *gin.RouterGroup {
	r := t.r.Group("")
	r.Use()
	return r
}
