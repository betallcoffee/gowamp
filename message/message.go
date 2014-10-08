// message.go
package message

import "strconv"

import ()

// message
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

// roler
const (
	PUBLISHER = 1 << iota
	SUBSCRIBER
	CALLER
	CALLEE
)

const (
	BROKER = 1 << iota
	DEALER
)

type Response interface {
	Which() (int64, *MsgError)
	Array() Message
}

type Message []interface{}

func (m Message) Which() (int64, *MsgError) {
	which, ok := m[0].(int64)
	if ok {
		return which, nil
	}
	return 0, NewMsgErrorDetail(INVALIDARGUMENT, "the type of first element in Message is not int64.")
}

func (m Message) Array() Message {
	return m
}

func (m Message) String() string {
	which, merr := m.Which()
	if merr != nil {
		return merr.Error()
	}
	return strconv.Itoa(int(which))
}

const (
	INVALIDURI             URL = "wamp.error.invalid_uri"
	NOSUCHPROCEDURE        URL = "wamp.error.no_such_procedure"
	PROCEDUREALREADYEXISTS URL = "wamp.error.procedure_already_exists"
	NOSUCHREGISTRATION     URL = "wamp.error.no_such_registration"
	NOSUCHSUBSCRIPTION     URL = "wamp.error.no_such_subscription"
	INVALIDARGUMENT        URL = "wamp.error.invalid_argument"
	SYSTEMSHUTDOWN         URL = "wamp.error.system_shutdown"
	CLOSEREALM             URL = "wamp.error.close_realm"
	GOODBYEANDOUT          URL = "wamp.error.goodbye_and_out"
	NOTAUTHORIZED          URL = "wamp.error.not_authorized"
	AUTHORIZATIONFAILED    URL = "wamp.error.authorization_failed"
	NOSUCHREALM            URL = "wamp.error.no_such_realm"
	NOSUCHROLE             URL = "wamp.error.no_such_role"
)

type URL string

func (u URL) invalid() bool {
	return true
}

func (u URL) string() string {
	return string(u)
}

type MsgError struct {
	Reason URL
	Detail map[string]interface{}
}

func NewMsgError(reason URL) *MsgError {
	e := &MsgError{
		Reason: reason,
	}
	return e
}

func NewMsgErrorDetail(reason URL, msg string) *MsgError {
	e := &MsgError{
		Reason: reason,
	}
	e.Detail["message"] = msg
	return e
}

func (e *MsgError) Error() string {
	if msg, ok := e.Detail["message"]; ok {
		if s, yes := msg.(string); yes {
			return e.Reason.string() + s
		}
		return e.Reason.string() + " detail message is not string"
	}
	return e.Reason.string()
}
