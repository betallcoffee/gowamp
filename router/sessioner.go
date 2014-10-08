// sessioner.go
package router

import (
	"gowamp/message"
)

var sessionID int64 = 0

type session struct {
	id      int64
	roles   int8
	waitbye bool
}

func Hello(h *message.Hello) (*message.Welcome, *session) {
	sessionID++
	s := &session{
		id:      sessionID,
		roles:   h.Roles,
		waitbye: false,
	}
	w := message.NewWelcome(s.id)
	w.SetRoles(message.BROKER | message.DEALER)
	return w, s
}

func Goodbye(s *session, g *message.Goodbye) *message.Goodbye {
	return message.NewGoodbye(message.GOODBYEANDOUT)
}
