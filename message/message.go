// message.go
package message

import (
	"encoding/json"
	"errors"
)

const (
	HELLO          = 1
	WELCOME        = 2
	ABORT          = 3
	CHALLENGE      = 4
	AUTHENTICATION = 5
	GOODBYE        = 6
	HEARTBEAT      = 7
	ERROR          = 8
)
const (
	PUBLISH   = 16
	PUBLISHED = 17
)
const (
	SUBSCRIBE    = 32
	SUBSCRIBED   = 33
	UNSUBSCRIBE  = 34
	UNSUBSCRIBED = 35
	EVENT        = 36
)
const (
	CALL   = 48
	CANCEL = 49
	RESULT = 50
)
const (
	REGISTER     = 64
	REGISTERED   = 65
	UNREGISTER   = 66
	UNREGISTERED = 67
	INVOCATION   = 68
	INTERRUPT    = 69
	YIELD        = 70
)

type Response interface {
	Array() *Message
}

type Message []interface{}

func (m *Message) Which() (int64, error) {
	which, ok := m[0].(int64)
	if ok {
		return which, nil
	}
	return nil, NewMsgerrorDetail(INVALIDARGUMENT, "the type of first element in Message is not int64.")
}

const (
	INVALIDURI             = "wamp.error.invalid_uri"
	NOSUCHPROCEDURE        = "wamp.error.no_such_procedure"
	PROCEDUREALREADYEXISTS = "wamp.error.procedure_already_exists"
	NOSUCHREGISTRATION     = "wamp.error.no_such_registration"
	NOSUCHSUBSCRIPTION     = "wamp.error.no_such_subscription"
	INVALIDARGUMENT        = "wamp.error.invalid_argument"
	SYSTEMSHUTDOWN         = "wamp.error.system_shutdown"
	CLOSEREALM             = "wamp.error.close_realm"
	GOODBYEANDOUT          = "wamp.error.goodbye_and_out"
	NOTAUTHORIZED          = "wamp.error.not_authorized"
	AUTHORIZATIONFAILED    = "wamp.error.authorization_failed"
	NOSUCHREALM            = "wamp.error.no_such_realm"
	NOSUCHROLE             = "wamp.error.no_such_role"
)

type URL string

func (u *URL) invalid() bool {
	return true
}

type msgerror struct {
	Reason URL
	Detail map[string]interface{}
}

func NewMsgerror(reason URL) error {
	return &msgerror{
		Reason: reason,
		Detail: {},
	}
}

func NewMsgerrorDetail(reason URL, msg string) error {
	return &msgerror{
		Reason: reason,
		Detail: {"message": msg},
	}
}

func (e *msgerror) Error() string {
	if msg, ok := e.Detail["message"]; ok {
		return e.Reason + e.Detail["message"]
	}
	return e.Reason
}
