// websocket.go
package transport

import (
	"github.com/gorilla/websocket"
	log "llog"
	"net/http"
	"time"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Websocket struct {
	conn *websocket.Conn
}

func NewWebsocket(w http.ResponseWriter, r *http.Request) (*Websocket, error) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &Websocket{
		conn: conn,
	}, nil
}

func (w *Websocket) WriteMessage(data []byte) error {
	return w.conn.WriteMessage(websocket.TextMessage, data)
}

func (w *Websocket) ReadMessage() ([]byte, error) {
	_, data, err := w.conn.ReadMessage()
	return data, err
}

func (w *Websocket) SetReadDeadline(t time.Time) error {
	return w.conn.SetReadDeadline(t)
}

func (w *Websocket) SetWriteDeadline(t time.Time) error {
	return w.conn.SetWriteDeadline(t)
}
