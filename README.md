# page_check

This is a small script I used to check if my web page is up and running.
It is run in a cron job every 10 minutes.

I use the [SendGrid](https://sendgrid.com) API to mail myself an alert if the
page is down. I figured this would be easier than setting up SMTP myself.

If you wish to run the script yourself, here are some instructions.

Install
* Go programming language: https://golang.org/doc/install
* SendGrid Go package: https://github.com/sendgrid/sendgrid-go
* SendGrid API key: https://sendgrid.com/beta/settings/api_key

You must set your SendGrid API key as an environment variable.
```sh
export SENDGRID_API_KEY=some_sample_123
```

Then you can build the binary using `go build`.

# License
MIT
