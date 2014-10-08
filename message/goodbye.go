// goodbye.go
package message

type Goodbye struct {
	which   int64
	Message string
	Details map[string]interface{}
	Reason  URL
}

func NewGoodbye(reason URL) *Goodbye {
	g := new(Goodbye)
	g.which = GOODBYE
	g.Reason = reason
	return g
}

func NewGoodbyeByMessage(m Message) (*Goodbye, *MsgError) {
	g := new(Goodbye)
	g.which, _ = m.Which()
	reason, ok := m[2].(URL)
	if !ok {
		return nil, NewMsgErrorDetail(INVALIDURI, "the reason of Goodbye message is not URL(string).")
	}
	if reason.invalid() {
		return nil, NewMsgErrorDetail(INVALIDURI, "the reason of Goodbye message is invlidation URL(string).")
	}
	g.Reason = reason

	merr := g.parseDetials(m)
	if merr != nil {
		return nil, merr
	}

	return g, nil
}

func (g *Goodbye) Which() (int64, *MsgError) {
	return g.which, nil
}

func (g *Goodbye) Array() Message {
	return []interface{}{g.which, g.Details, g.Reason}
}

func (g *Goodbye) parseDetials(m Message) *MsgError {
	details, ok := m[1].(map[string]interface{})
	if !ok {
		return NewMsgErrorDetail(INVALIDARGUMENT, "the details of Goodbye message is not map.")
	}

	message, ok := details["message"]
	if !ok {
		return NewMsgErrorDetail(INVALIDARGUMENT, "the details of Goodbye message do not contain message.")
	}

	g.Message, ok = message.(string)
	if !ok {
		return NewMsgErrorDetail(INVALIDARGUMENT, "the message of Goodbye message details is not string.")
	}
	return nil
}
