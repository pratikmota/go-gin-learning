package main

import (
	"fmt"
	"net/http"

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
			"message": "pong语言",
		})
	})

	//#2 using function
	r.GET("/hello2", hello2)

	//#3 Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.
	r.GET("/AsciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	//https://sourjp.github.io/posts/go-gin/
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
