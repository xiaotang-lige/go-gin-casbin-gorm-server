package route

import (
	"github.com/gin-gonic/gin"
	"messageServe/server"
)

func (t *route) loggin(r *gin.RouterGroup) {
	r1 := r.Group("login")
	r1.POST("query", server.ServerApi.QueryAccountAndPassword)
}
