package route

import (
	"github.com/gin-gonic/gin"
	"messageServe/server"
)

func (t *route) userConfig(r *gin.RouterGroup) {
	api := r.Group("userConfig")
	api.GET("query", server.ServerApi.UserConfig.UserConfig)
}
