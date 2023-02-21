package engine

import (
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strings"

	"github.com/shahariaazam/mig/pkg/message"
)

type SMTP struct {
	Username string
	Password string
	Host     string
	Port     string
}

func NewSMTP(username, password, host, port string) *SMTP {
	return &SMTP{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func (s *SMTP) Send(msg message.Message) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	toForMailHeader := make([]string, len(msg.To))
	toWithoutName := make([]string, len(msg.To))
	for i, recipient := range msg.To {
		toForMailHeader[i] = recipient.String()
		toWithoutName[i] = recipient.Address
	}
	headers := make(map[string]string)
	headers["From"] = msg.From.String()
	headers["To"] = strings.Join(toForMailHeader, ", ")
	headers["Subject"] = msg.Subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=utf-8"
	headers["Content-Transfer-Encoding"] = "quoted-printable"
	var body strings.Builder
	quotedPrintableWriter := quotedprintable.NewWriter(&body)
	quotedPrintableWriter.Write([]byte(msg.Text))
	quotedPrintableWriter.Close()
	m := buildMessage(headers, body.String())
	err := smtp.SendMail(s.Host+":"+s.Port, auth, msg.From.Address, toWithoutName, []byte(m))
	if err != nil {
		return err
	}
	return nil
}

func buildMessage(headers map[string]string, body string) string {
	var msg strings.Builder
	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n")
	msg.WriteString(body)
	return msg.String()
}
