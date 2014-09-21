// hello.go
package message

import (
	"encoding/json"
)

const (
	PUBLISHER = 1 << iota
	SUBSCRIBER
	CALLER
	CALLEE
)

type hello struct {
	Which int64
	Realm URL
	roles int8
}

func NewHelloByMessage(m *Message) (*hello, error) {
	h := new(hello)
	h.Which = m.Which()
	realm, ok := m[1].(URL)
	if !ok {
		return nil, NewMsgerrorDetail(INVALIDURI, "the realm of Hello message is not URL(string).")
	}
	if realm.invalid() {
		return nil, NewMsgerrorDetail(INVALIDURI, "the realm of Hello message is invlidation URL(string).")
	}
	h.Realm = realm

	err := h.parseDetails(m)
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *hello) parseDetails(m *Message) error {
	h.roles = 0
	details, ok := m[2].(map[string]interface{})
	if !ok {
		return NewMsgerrorDetail(INVALIDARGUMENT, "the details of Hello message is not map.")
	}
	roles_, ok := details["roles"]
	if !ok {
		return NewMsgerrorDetail(INVALIDARGUMENT, "the details of Hello message do not contain roles.")
	}
	roles, ok := roles_.(map[string]interface{})
	if !ok {
		return NewMsgerrorDetail(INVALIDARGUMENT, "the roles of Hello message is not map.")
	}

	publisher, ok := roles["publisher"]
	if ok {
		h.roles |= PUBLISHER
	}
	subscriber, ok := roles["subscriber"]
	if ok {
		h.roles |= SUBSCRIBER
	}
	caller, ok := roles["caller"]
	if ok {
		h.roles |= CALLER
	}
	callee, ok := roles["callee"]
	if ok {
		h.roles |= CALLEE
	}
	if h.roles == 0 {
		return NewMsgerrorDetail(NOSUCHROLE, "the roles of Hello message is empty.")
	}
	return nil
}
