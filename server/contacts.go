package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"messageServe/dao"
	"messageServe/model"
	"messageServe/tool"
)

type contacts struct {
}

func (*contacts) ShowAll(c *gin.Context) {
	user := &model.UserConfig{UserId: tool.Api.Token.GetClaimsModel(c).UserId}
	listData, db := dao.Api.Contacts.QueryAll(&model.Contacts{A: user.UserId, State: 1})
	if db.Error != nil {
		log.Println(db.Error)
		tool.Api.Response.Write(c, &model.Response{State: 403, Err: "获取全部联系人错误！"})
	}
	tool.Api.Response.Write(c, &model.Response{Len: int(db.RowsAffected), State: 200, Data: listData})
}
