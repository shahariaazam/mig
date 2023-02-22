package engine

import (
	"errors"
	"github.com/shahariaazam/mig/pkg/message"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/mail"
	"strings"
	"testing"

	"github.com/emersion/go-smtp"
)

// The Backend implements SMTP server methods.
type Backend struct {
	receivedData string
}

func (bkd *Backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	return &Session{backend: bkd}, nil
}

// A Session is returned after EHLO.
type Session struct {
	backend *Backend
}

func (s *Session) AuthPlain(username, password string) error {
	if username != "username" || password != "password" {
		return errors.New("Invalid username or password")
	}
	return nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		s.backend.receivedData = string(b)
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func (bkd *Backend) GetReceivedData() string {
	return bkd.receivedData
}

func TestSend(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	backend := &Backend{}
	s := smtp.NewServer(backend)
	s.Domain = "localhost"
	s.AllowInsecureAuth = true
	go s.Serve(l)

	_, err = net.Dial("tcp", l.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	smtpServer := strings.Split(l.Addr().String(), ":")
	smtpClient := NewSMTP("username", "password", smtpServer[0], smtpServer[1])

	msg := message.Message{
		From: mail.Address{
			Name:    "John Doe",
			Address: "johndoe@example.com",
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

	err = smtpClient.Send(msg)
	assert.NoError(t, err)

	// Get the received data from the backend
	receivedData := backend.GetReceivedData()

	m, err := mail.ReadMessage(strings.NewReader(receivedData))
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "text/plain; charset=utf-8", m.Header.Get("Content-Type"))
	assert.Equal(t, "quoted-printable", m.Header.Get("Content-Transfer-Encoding"))
	assert.Equal(t, "\"John Doe\" <johndoe@example.com>", m.Header.Get("From"))
	assert.Equal(t, "\"Jane Smith\" <janesmith@example.com>", m.Header.Get("To"))
	assert.Equal(t, "1.0", m.Header.Get("MIME-Version"))
	body, err := ioutil.ReadAll(m.Body)
	assert.Equal(t, "This is a test email\r\n", string(body))
}
