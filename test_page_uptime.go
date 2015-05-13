package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

func sendMail() {
	api_key := os.Getenv("SENDGRID_API_KEY")
	if api_key == "" {
		fmt.Println("Could not find the SendGrid API key!")
		os.Exit(1)
	}

	sg := sendgrid.NewSendGridClientWithApiKey(api_key)
	message := sendgrid.NewMail()
	message.AddTo("you@youremail.com")
	message.SetSubject("Alert!")
	message.SetText("Could not reach web page. Fix it!")
	message.SetFrom("you@youremail.com")

	r := sg.Send(message)
	if r == nil {
		fmt.Println("Email sent!")
	} else {
		fmt.Println(r)
	}
}

func checkSite() {
	resp, err := http.Head("somewebsite.com")

	if err != nil {
		fmt.Println("Could not perform HEAD request!")
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		sendMail()
		fmt.Println("Sent email! Not good!")
	} else {
		fmt.Println("Got a 200, everything is fine")
	}
}

func main() {
	checkSite()
}
