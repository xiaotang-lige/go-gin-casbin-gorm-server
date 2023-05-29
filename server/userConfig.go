package server

import (
	"github.com/gin-gonic/gin"
	"messageServe/dao"
	"messageServe/model"
	"messageServe/tool"
)

type UserConfig struct {
}

func (t *UserConfig) Login(c *gin.Context) {
	config := &model.UserConfig{}
	err := c.BindJSON(config)
	if err != nil {
		tool.Api.Response.Write(c, &model.Response{State: 400, Err: "参数错误！"})
		return
	}
	db, userData := dao.Api.UserConfig.Query(config)
	if db.RowsAffected == 0 {
		tool.Api.Response.Write(c, &model.Response{State: 400, Err: "账号密码错误！"})
		return
	}
	userData.Token = tool.Api.Token.Issue(userData)
	tool.Api.Response.Write(c, &model.Response{State: 200, Len: int(db.RowsAffected), Data: userData})
}
