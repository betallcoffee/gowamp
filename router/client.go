// client.go
package router

import (
	"errors"
	log "llog"
	"message"
	"serialization"
	"strconv"
	"transport"
)

type CliError struct {
	C   *client
	Err error
}

func (e *CliError) Error() string {
	if e == nil {
		return "<nil>"
	}
	s := strconv.Itoa(e.c.sessionID)
	s += " req" + strconv.Itoa(e.c.reqCount)
	if e.c.addr != nil {
		s += " " + e.c.addr
	}
	s += ": " + e.Err.Error()
	return s
}

type client struct {
	reqCount  int64
	addr      string
	sessionID int64
	roles     int8
	waitBye   bool

	transport     *transport.Transport
	serialization *serialization.Serialization
	dealer        *router.Dealer
	broker        *router.Broker
}

func NewClient(transport *transport.Transport, serialization *serialization.Serialization,
	dealer *router.Dealer, broker *router.Broker) *client {
	return &client{
		reqCount:  0,
		andr:      "",
		sessionID: 0,
		roles:     0,
		waitBye:   false,

		transport:     transport,
		serialization: serialization,
		dealer:        dealer,
		broker:        broker,
	}
}

func (c *client) inroutine() {
	for {
		c.reqCount++
		var response *message.Response

		msg, err := c.readMessage()
		if err != nil {
			log.Error(CliError(c, err))
			response = message.NewAbort()
		} else {
			response = c.dispatchMessage(msg)
		}

		c.writeMessage(response.array())
	}
}

func (c *client) readMessage() (*message.Message, error) {
	buffer, err := c.transport.ReadMessage()
	if err != nil {
		log.Error(CliError(c, err))
		return nil, err
	}
	datas, err := c.serialization.Decode(buffer)
	if err != nil {
		log.Error(CliError(c, err))
		return nil, err
	}
	msg, ok := datas.(message.Message)
	if !ok {
		return nil, errors.New("the message format is error.")
	}
	return &msg, nil
}

func (c *client) dispatchMessage(msg *message.Message) *message.Message {
	which, err := msg.Which()
	if err != nil {
		log.Error(CliError(c, err))
		return message.NewAbort()
	}

	switch which {
	case message.HELLO:
		h, err := message.NewHelloByMessage(msg)
		if err != nil {
			log.Error(CliError(c, err))
		}
	}
}

func (c *client) outroutine() {

}

func (c *client) writeMessage(msg *message.Message) {

}
