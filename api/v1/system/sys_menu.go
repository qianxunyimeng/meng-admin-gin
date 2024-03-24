package system

import (
	"github.com/gin-gonic/gin"
)

type MenuApi struct{}

func (a *MenuApi) GetMenu(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "menu",
	})
}
