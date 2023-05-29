package route

import (
	"github.com/gin-gonic/gin"
	"messageServe/server"
)

func (*route) contacts(r *gin.RouterGroup) {
	t := r.Group("linkman")
	t.GET("all", server.Api.Contacts.ShowAll)
}
