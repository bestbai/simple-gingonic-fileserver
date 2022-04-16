package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	c.HTML(http.StatusOK, "select_file.html", gin.H{})
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
