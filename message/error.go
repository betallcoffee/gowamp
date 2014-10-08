// error.go
package message

type Error struct {
	which    int64
	reqWhich int64
	reqID    int64
	details  map[string]interface{}
	url      URL
	args     []interface{}
	argsKw   map[string]interface{}
}

func NewError(reqWhich, reqID int64, details map[string]interface{},
	url URL, args []interface{}, argsKw map[string]interface{}) *Error {
	return &Error{
		which:    ERROR,
		reqWhich: reqWhich,
		reqID:    reqID,
		details:  details,
		url:      url,
		args:     args,
		argsKw:   argsKw,
	}
}

func (e *Error) Which() (int64, *MsgError) {
	return e.which, nil
}
func (e *Error) Array() Message {
	return []interface{}{e.which, e.reqWhich, e.reqID, e.details, e.url, e.args, e.argsKw}
}
