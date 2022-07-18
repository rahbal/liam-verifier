package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"

	"github.com/heroku/go-getting-started/code"
)

func main() {
	port := os.Getenv("PORT")
	result := code.Email_result{}
	var err error

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

		emails := c.PostForm("emails")

		result, err = code.Verifier(emails)
		// log.Output(1, message+":"+err.Error())

		// fmt.Println("main went ahead")

		if err != nil {
			return
		}

		c.HTML(http.StatusOK, "result.tmpl.html", gin.H{
			"mx_records": result.Mx_slice,
			"email":      result.Email_add,
		})

	})

	router.Run(":" + port)
}
