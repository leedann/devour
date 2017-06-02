package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/leedann/devour/devoursvr/sessions"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//WebSocketUpgradeHandler ensures that the user is authenticated, upgrades the client and sends it to the notifier
func (ctx *Context) WebSocketUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	//ensure that the user is authenticated
	// https://golang.org/pkg/net/url/#URL.Query
	auth := r.URL.Query().Get("auth")

	s, err := sessions.GetSessionID(r, ctx.SessionKey)
	if err != nil {
		log.Println("could not find id")
		return
	}

	if s.String() == auth {
		//user the gorilla upgrader
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		//pass upgraded to notifier
		ctx.Notifier.AddClient(conn)
	} else {
		log.Println("User not authenticated")
		return
	}
}
