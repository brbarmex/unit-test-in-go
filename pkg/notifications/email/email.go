package email

import "fmt"

func SendMail(body string, to string) {
	fmt.Println("Sending mail to:" + to + body)
}
