package route

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"messageServe/tool"
	"os"
)

type route struct {
	r *gin.Engine
}

var RouteApi = new(route)

func Main() {
	RouteApi.link()
	//需要权限
	{
		RouteApi.contacts(RouteApi.control())
	}
	//无需权限
	{
		RouteApi.userConfig(RouteApi.open())
	}
	err := RouteApi.r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
func (t *route) link() {
	f, _ := os.Create(tool.ProjectPath() + "/route/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	t.r = gin.Default()
}

func (t *route) control() *gin.RouterGroup {
	r := t.r.Group("")
	r.Use(tool.TokenHande()).Use(tool.CasbinHandle())
	return r
}

func (t *route) open() *gin.RouterGroup {
	r := t.r.Group("")
	r.Use()
	return r
}
