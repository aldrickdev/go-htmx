package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Setup(c *gin.Context) {
	fmt.Println("In this page because we need a github pat")
	c.HTML(200, "index", nil)
}
