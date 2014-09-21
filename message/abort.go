// abort.go
package message

type abort struct {
	which   int64
	details map[string]interface{}
	reason  URL
}

func NewAbort(err msgerror) *abort {
	return &abort{
		which:   ABORT,
		details: err.Detail,
		reason:  err.Reason,
	}
}

func (a *abort) Array() *Message {
	return []interface{}{which, details, reason}
}
