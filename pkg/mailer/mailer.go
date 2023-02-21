package mailer

import (
	"github.com/shahariaazam/mig/pkg/engine"
	"github.com/shahariaazam/mig/pkg/message"
)

type Mailer struct {
	engine engine.Engine
}

func NewMailer(engine engine.Engine) Mailer {
	return Mailer{engine: engine}
}

func (m Mailer) Deliver(message message.Message) error {
	return m.engine.Send(message)
}
