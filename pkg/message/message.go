package message

import (
	"net/mail"
)

type Message struct {
	To      []mail.Address
	From    mail.Address
	Subject string
	Text    string
	Html    string
}
