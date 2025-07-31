package api

import (
	"github.com/Samrat-Mukherjee/file_share/internals/api/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/upload-file", handlers.HandleFile)

	return router

}
