package code

import (
	"fmt"
	"net"

	emailverifier "github.com/AfterShip/email-verifier"

	_ "github.com/heroku/x/hmetrics/onload"
)

type Email_result struct {
	Mx_slice  []*net.MX
	Email_add string
}

var (
	verifier = emailverifier.NewVerifier()
)

func Verifier(mail_string string) (Email_result, error) {

	fmt.Print(mail_string)
	// email := "example@exampledomain.org"
	email := mail_string
	var e_result Email_result
	e_result.Email_add = mail_string

	// parse email
	syntax := verifier.ParseAddress(mail_string)

	// check and print MC records
	mx, err := verifier.CheckMX(syntax.Domain)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return e_result, err
	}
	if mx.HasMXRecord {
		fmt.Println("Detail about MX Records of Domain")
		fmt.Println(mx.Records)
		e_result.Mx_slice = mx.Records
		for k, v := range mx.Records {
			fmt.Print(k, v)
		}
	}

	// verify bunch of email things
	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return e_result, err
	}
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return e_result, err
	}

	fmt.Println("email validation result: ", *ret)

	return e_result, err

}
