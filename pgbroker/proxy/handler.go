package proxy

import (
	"github.com/rueian/pgbroker/message"
)

type MessageHandler func(*Ctx, []byte) (message.Reader, error)
type MessageHandlerOther func(*Ctx, byte, []byte) (message.Reader, error)

type MessageHandlerRegister interface {
	GetHandler(byte) MessageHandler
}
