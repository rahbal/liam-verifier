package code

import (
	"fmt"

	_ "github.com/heroku/x/hmetrics/onload"
)

func Verifier(mail_string string) string {

	fmt.Print(mail_string)

	return mail_string

}
