package src

import (
	"simple_fileserver/src/pages"

	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {
	r.GET("/", pages.Index)
	r.GET("/hello", pages.Hello)
	r.POST("/upload", pages.Upload)
}
