package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// code.Verifier() // function name need to be Capital to export

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/result", func(c *gin.Context) {

		message := c.PostForm("msg")
		fmt.Print(message)

		c.HTML(http.StatusOK, "result.tmpl.html", gin.H{
			"message": message,
		})

	})

	router.Run(":" + port)
}
