package code

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func Verifier(c *gin.Context) {
	fmt.Print("in verifier")

	c.HTML(http.StatusOK, "result.tmpl.html", nil)

}
