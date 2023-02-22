package engine

import "github.com/shahariaazam/mig/pkg/message"

type Engine interface {
	Send(message message.Message) error
}
