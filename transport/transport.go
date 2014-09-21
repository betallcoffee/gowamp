// transport.go
package transport

type Transport interface {
	WriteMessage([]byte) error
	ReadMessage() ([]byte, error)
}
