// gowamp.go
package main

import (
	"net/http"
	"os"
	"gowamp/transport"
	log "llog"
	"gowamp/router"
	"gowamp/serialization"
	"sync"
)

var clients map[string]*router.Client
var mutex sync.Mutex

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := transport.NewWebsocket(w, r)
	if err != nil {
		log.Error(err)
		return
	}
	c := router.NewClient(t, serialization.NewJSON(), router.NewDealer(), router.NewBroker())
	mutex.Lock()
	clients[r.RemoteAddr] = c
	mutex.Unlock()
	go c.Inroutine()
	c.Outroutine()
	mutex.Lock()
	delete(clients, r.RemoteAddr)
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
	os.Exit(0)
}
