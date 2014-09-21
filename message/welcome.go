// welcome.go
package message

const (
	BROKER = 1 << iota
	DEALER
)

type welcome struct {
	which     int64
	sessionID int64
	details   map[string]interface{}

	roles int8
}

func NewWelcome(sessionID int64) {
	return &welcome{
		which:     WELCOME,
		sessionID: sessionID,
		roles:     0,
		details:   {"roles": {}},
	}
}

func (w *welcome) Array() *Message {
	return []interface{}{w.which, w.sessionID, w.details}
}

func (w *welcome) SetRoles(roles int8) {
	w.roles = roles
	w.details = map[string]interface{}{"roles": {}}
	if roles & BROKER {
		w.details["roles"]["broker"] = map[string]interface{}{}
	}
	if roles & DEALER {
		w.details["roles"]["dealer"] = map[string]interface{}{}
	}
}

func (w *welcome) AddRoles(roles int8) {
	if roles & BROKER {
		w.roles |= BROKER
		w.details["roles"]["broker"] = map[string]interface{}{}
	}
	if roles & DEALER {
		w.roles |= DEALER
		w.details["roles"]["dealer"] = map[string]interface{}{}
	}
}

func (w *welcome) DelRoles(roles int8) {
	if roles & BROKER {
		w.roles &= ^BROKER
		delete(w.details["roles"], "broker")
	}
	if roles & DEALER {
		w.roles &= ^DEALER
		delete(w.details["roles"], "dealer")
	}
}
