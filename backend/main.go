package main

import (
	"simple_fileserver/src"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.LoadHTMLGlob("template/*")
	// r.StaticFS("/file", http.Dir("public"))

	src.GetRoute(r)
	r.Run(":8090")
}
