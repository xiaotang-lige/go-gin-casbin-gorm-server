package tool

import (
	"github.com/gin-gonic/gin"
	"messageServe/model"
)

type response struct {
}

func (*response) Write(c *gin.Context, model *model.Response) {
	c.JSON(model.State, model)
}
