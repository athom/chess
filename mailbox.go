package chess

type MailBox interface {
	Receive() *PlayerState
	Send(*GameState)
}
