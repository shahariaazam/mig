MIG is a Go library that enables you to send emails using various mail providers, including SMTP. It aims to make sending emails easy and intuitive, with a clean API and clear documentation.

Features
--------

-   Send emails using various mail providers, including SMTP.

## Installation

To install Mig, use `go get`:

```sh
go get github.com/shahariaazam/mig
```

Usage
-----

The library is very easy to use. Here's an example that shows how to send an email using SMTP:

```go
package main

import (
	"github.com/shahariaazam/mig/pkg/engine"
	"github.com/shahariaazam/mig/pkg/message"
	"net/mail"
)

func main() {
	// Create a new SMTP client
	smtpClient := engine.NewSMTP("username", "password", "smtp.example.com", "587")

	// Create a test message
	msg := message.Message{
		From: mail.Address{
			Name:    "Jane Smith",
			Address: "janesmith@example.com",
		},
		To: []mail.Address{
			{
				Name:    "Jane Smith",
				Address: "janesmith@example.com",
			},
		},
		Subject: "Test Email",
		Text:    "This is a test email",
	}

	// Send the email
	err := smtpClient.Send(msg)
	if err != nil {
		panic(err)
	}
}
```

Contributing
------------

Contributions are welcome. If you encounter any issues or have any feature requests, please create an issue on [GitHub](https://github.com/shahariaazam/mig).

License
-------

MIG is released under the [MIT license](https://github.com/shahariaazam/mig/blob/main/LICENSE).