package pages

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": string(err.Error()), "success": false})
	}

	filename := header.Filename
	out, err := os.Create("/api/public/" + filename)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": string(err.Error()), "success": false})
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": string(err.Error()), "success": false})
	}

	filepath := "http://192.168.18.50:8001/" + filename
	c.JSON(http.StatusOK, gin.H{"message": "file upload successful", "filepath": filepath, "success": true})
}
