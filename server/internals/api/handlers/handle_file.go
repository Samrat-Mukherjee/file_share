package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFile(c *gin.Context) {
	file, err := c.FormFile("fufu")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no file is received",
		})
		return
	}

	fmt.Println("File Name ", file.Filename)

	c.SaveUploadedFile(file, "./file"+file.Filename)

	c.JSON(http.StatusOK, gin.H{
		"Hello from Samrat ": file.Filename,
	})
}
