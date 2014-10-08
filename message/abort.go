// abort.go
package message

type Abort struct {
	which   int64
	details map[string]interface{}
	reason  URL
}

func NewAbort(err *MsgError) *Abort {
	return &Abort{
		which:   ABORT,
		details: err.Detail,
		reason:  err.Reason,
	}
}

func (a *Abort) Which() (int64, *MsgError) {
	return a.which, nil
}

func (a *Abort) Array() Message {
	return []interface{}{a.which, a.details, a.reason}
}
