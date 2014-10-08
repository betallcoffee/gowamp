// hello.go
package message

import (
)

type Hello struct {
	Which int64
	Realm URL
	Roles int8
}

func NewHelloByMessage(m Message) (*Hello, *MsgError) {
	h := new(Hello)
	h.Which, _ = m.Which()
	realm, ok := m[1].(URL)
	if !ok {
		return nil, NewMsgErrorDetail(INVALIDURI, "the realm of Hello message is not URL(string).")
	}
	if realm.invalid() {
		return nil, NewMsgErrorDetail(INVALIDURI, "the realm of Hello message is invlidation URL(string).")
	}
	h.Realm = realm

	merr := h.parseDetails(m)
	if merr != nil {
		return nil, merr
	}

	return h, nil
}

func (h *Hello) parseDetails(m Message) *MsgError {
	h.Roles = 0
	details, ok := m[2].(map[string]interface{})
	if !ok {
		return NewMsgErrorDetail(INVALIDARGUMENT, "the details of Hello message is not map.")
	}
	roles_, ok := details["roles"]
	if !ok {
		return NewMsgErrorDetail(INVALIDARGUMENT, "the details of Hello message do not contain roles.")
	}
	roles, ok := roles_.(map[string]interface{})
	if !ok {
		return NewMsgErrorDetail(INVALIDARGUMENT, "the roles of Hello message is not map.")
	}

	_, ok = roles["publisher"]
	if ok {
		h.Roles |= PUBLISHER
	}
	_, ok = roles["subscriber"]
	if ok {
		h.Roles |= SUBSCRIBER
	}
	_, ok = roles["caller"]
	if ok {
		h.Roles |= CALLER
	}
	_, ok = roles["callee"]
	if ok {
		h.Roles |= CALLEE
	}
	if h.Roles == 0 {
		return NewMsgErrorDetail(NOSUCHROLE, "the roles of Hello message is empty.")
	}
	return nil
}
