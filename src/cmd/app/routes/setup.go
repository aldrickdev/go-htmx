package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup(c *gin.Context) {
	c.HTML(200, "setup", nil)
}
