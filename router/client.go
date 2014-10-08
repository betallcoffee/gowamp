// client.go
package router

import (
	"gowamp/message"
	"gowamp/serialization"
	"gowamp/transport"
	log "llog"
	"strconv"
)

type CliError struct {
	C   *Client
	Err string
}

func NewCliError(c *Client, err string) *CliError {
	return &CliError{
		C: c,
		Err: err,
	}
}
func (e *CliError) Error() string {
	if e == nil {
		return "<nil>"
	}
	s := strconv.Itoa(int(e.C.session.id))
	s += " req" + strconv.Itoa(int(e.C.reqCount))
	s += " " + e.C.addr
	s += ": " + e.Err
	return s
}

type Client struct {
	reqCount  int64
	addr      string
	session   *session
	quitChan  chan bool
	inChan  chan bool
	outChan chan bool

	transport     transport.Transport
	serialization serialization.Serialization
	dealer        *Dealer
	broker        *Broker
}

func NewClient(transport transport.Transport, serialization serialization.Serialization,
	dealer *Dealer, broker *Broker) *Client {
	return &Client{
		reqCount:      0,
		addr:          "",
		session:       nil,
		transport:     transport,
		serialization: serialization,
		dealer:        dealer,
		broker:        broker,
	}
}

func (c *Client) Quit() {
	c.quitChan <- true
	<- c.inChan
	<- c.outChan
}

func (c *Client) Inroutine() {
	log.Trace(c.addr + " inroutine begin")
	defer log.Trace(c.addr + " inroutine end")

	for {
		c.reqCount++
		var response message.Response

		msg, merr := c.readMessage()
		if merr != nil {
			log.Error(NewCliError(c, merr.Error()))
			// TODO now abort the session when the message the invalid, but it too ugly.
			// should send the ERROR message most conditions.
			response = message.NewAbort(merr)
		} else {
			response = c.dispatchMessage(msg)
		}

		c.writeMessage(response.Array())
		which, _ := response.Which()
		switch which {
		case message.ABORT:
			// TODO if send a abort message, then close the connection.
			return
		case message.GOODBYE:
			return
		}
	}
}

func (c *Client) readMessage() (message.Message, *message.MsgError) {
	log.Trace(c.addr + " readMessage begin")
	defer log.Trace(c.addr + " readMessage end")

	buffer, err := c.transport.ReadMessage()
	if err != nil {
		log.Error(NewCliError(c, err.Error()))
		return nil, message.NewMsgErrorDetail(message.INVALIDARGUMENT, err.Error())
	}
	datas, err := c.serialization.Decode(buffer)
	if err != nil {
		log.Error(NewCliError(c, err.Error()))
		return nil, message.NewMsgErrorDetail(message.INVALIDARGUMENT, err.Error())
	}
	msg := message.Message(datas)
//	if !ok {
//		return nil, message.NewMsgErrorDetail(message.INVALIDARGUMENT, "the message format is error.")
//	}
	log.Debug(c.addr + " readMessage " + msg.String())
	return msg, nil
}

func (c *Client) dispatchMessage(msg message.Message) message.Response {
	log.Trace(c.addr + " dispatchMessage begin " + msg.String())
	defer log.Trace(c.addr + " dispatchMessage end " + msg.String())

	which, merr := msg.Which()
	if merr != nil {
		log.Error(NewCliError(c, merr.Error()))
		// TODO now abort the session when the message the invalid, but it too ugly.
		// should send the ERROR message most conditions.
		return message.NewAbort(merr)
	}

	switch which {
	case message.HELLO:
		h, merr := message.NewHelloByMessage(msg)
		if merr != nil {
			log.Error(NewCliError(c, merr.Error()))
			return message.NewAbort(merr)
		}
		w, session := Hello(h)
		c.session = session
		return w
	case message.GOODBYE:
		g1, merr := message.NewGoodbyeByMessage(msg)
		if merr != nil {
			log.Error(NewCliError(c, merr.Error()))
			return message.NewAbort(merr)
		}
		g2 := Goodbye(c.session, g1)
		return g2
	}
	return nil
}

func (c *Client) Outroutine() {

}

func (c *Client) writeMessage(msg message.Message) {

}
