package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"messageServe/dao"
	"messageServe/model"
)

type UserConfig struct {
}

func (t *UserConfig) UserConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "123",
	})
}
func (t *UserConfig) QueryAccountAndPassword(c *gin.Context) {
	config := &model.UserConfig{}
	err := c.BindJSON(config)
	if err != nil {
		log.Println(err)
		return
	}
	dao.DaoApi.Create(config)
}
