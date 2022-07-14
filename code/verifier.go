package code

import (
	"fmt"

	_ "github.com/heroku/x/hmetrics/onload"
)

func verifier() {
	fmt.Print("in verifier")
}
