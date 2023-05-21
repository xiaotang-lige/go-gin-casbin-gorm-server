package server

import "github.com/gin-gonic/gin"

type UserConfig struct {
}

func (t *UserConfig) UserConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "123",
	})
}
