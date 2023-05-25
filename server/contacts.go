package server

import (
	"github.com/gin-gonic/gin"
	"messageServe/dao"
	"messageServe/model"
	"messageServe/tool"
)

type contacts struct {
}

func (*contacts) ShowAll(c *gin.Context) {
	user := &model.UserConfig{UserId: tool.Api.Token.GetClaimsModel(c).UserId}
	listData, db := dao.Api.Contacts.QueryAll(user, 1)
	if db.Error != nil {
		tool.Api.Response.Write(c, &model.Response{State: 403, Err: "获取全部联系人错误！"})
	}
	tool.Api.Response.Write(c, &model.Response{
		Len:   int(db.RowsAffected),
		State: 200,
		Err:   "",
		Data:  listData,
	})
}
