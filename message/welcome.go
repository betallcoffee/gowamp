// welcome.go
package message

type Welcome struct {
	which     int64
	sessionID int64
	details   map[string]map[string]interface{}

	roles int8
}

func NewWelcome(sessionID int64) *Welcome {
	w := &Welcome{
		which:     WELCOME,
		sessionID: sessionID,
		roles:     0,
	}
	w.details["roles"] = make(map[string]interface{}, 0)
	return w
}

func (w *Welcome) Which() (int64, *MsgError) {
	return w.which, nil
}

func (w *Welcome) Array() Message {
	return []interface{}{w.which, w.sessionID, w.details}
}

func (w *Welcome) SetRoles(roles int8) {
	w.roles = roles
	w.details["roles"] = make(map[string]interface{}, 0)
	if roles & BROKER == 0 {
		w.details["roles"]["broker"] = make(map[string]interface{}, 0)
	}
	if roles & DEALER == 0 {
		w.details["roles"]["dealer"] = make(map[string]interface{}, 0)
	}
}

func (w *Welcome) AddRoles(roles int8) {
	if roles & BROKER == 0 {
		w.roles |= BROKER
		w.details["roles"]["broker"] = make(map[string]interface{}, 0)
	}
	if roles & DEALER == 0 {
		w.roles |= DEALER
		w.details["roles"]["dealer"] = make(map[string]interface{}, 0)
	}
}

func (w *Welcome) DelRoles(roles int8) {
	if roles & BROKER == 0 {
		w.roles &= ^BROKER
		delete(w.details["roles"], "broker")
	}
	if roles & DEALER == 0 {
		w.roles &= ^DEALER
		delete(w.details["roles"], "dealer")
	}
}
