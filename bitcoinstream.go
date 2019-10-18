package main

import (
	"log"
	"time"

	"github.com/rgamba/evtwebsocket"
)

func main() {
	c := evtwebsocket.Conn{

		// When connection is established
		OnConnected: func(w *evtwebsocket.Conn) {
			log.Println("Connected")
		},

		// When a message arrives
		OnMessage: func(msg []byte, w *evtwebsocket.Conn) {
			log.Printf("OnMessage: %s\n", msg)
		},

		// When the client disconnects for any reason
		OnError: func(err error) {
			log.Printf("** ERROR **\n%s\n", err.Error())
		},

		// This is used to match the request and response messagesP>termina
		MatchMsg: func(req, resp []byte) bool {
			return string(req) == string(resp)
		},

		// Auto reconnect on error
		Reconnect: true,

		// Set the ping interval (optional)
		PingIntervalSecs: 5,

		// Set the ping message (optional)
		PingMsg: []byte("PING"),
	}

	// Connect
	err := c.Dial("wss://ws.blockchain.info/inv", "")
	if err != nil {
		log.Fatal(err)
	}

	outmsg := evtwebsocket.Msg{
		Body: []byte("{\"op\":\"unconfirmed_sub\"}"),
		Callback: func(resp []byte, w *evtwebsocket.Conn) {
			log.Printf("[%d] Callback: %s\n", 0, resp)
		},
	}

	//err = c.Send(msg)
	if err := c.Send(outmsg); err != nil {
		log.Println("Unable to send: ", err.Error())
	}

	time.Sleep(time.Second * 20)
}