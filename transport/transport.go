// transport.go
package transport

import "time"

type Transport interface {
	WriteMessage([]byte) error
	ReadMessage() ([]byte, error)
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}
