package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (c *Config) Start() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", getAllBooks)
		v1.POST("/books", createBook)
		v1.GET("/books/:id", getBookById)
		v1.POST("/books/:id/update", updateBookById)
		v1.POST("/books/:id/delete", deleteBookById)
	}

	listenPort := fmt.Sprintf(":%s", c.ListenPort)
	router.Run(listenPort)
}
