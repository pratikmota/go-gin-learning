package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func hello2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong2",
	})
}

func main() {
	fmt.Println("hello")

	r := gin.Default()
	//#1
	// GET: 127.0.0.1:8080/hello
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//#2
	r.GET("/hello2", hello2)

	r.Run() // listen and serve on 0.0.0.0:8080
}
