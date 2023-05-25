package route

import (
	"github.com/gin-gonic/gin"
	"messageServe/server"
)

func (t *route) userConfig(r *gin.RouterGroup) {
	api := r.Group("login")
	api.POST("query", server.Api.UserConfig.Login)
}
