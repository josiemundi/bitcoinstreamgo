package main

import (
	"encoding/json"
	"fmt"
	"github.com/rgamba/evtwebsocket"
	"log"
)

type Transaction struct {
	Op string `json:"op"`
	X  struct {
		LockTime int `json:"lock_time"`
		Ver      int `json:"ver"`
		Size     int `json:"size"`
		Inputs   []struct {
			Sequence int64 `json:"sequence"`
			PrevOut  struct {
				Spent   bool   `json:"spent"`
				TxIndex int    `json:"tx_index"`
				Type    int    `json:"type"`
				Addr    string `json:"addr"`
				Value   int    `json:"value"`
				N       int    `json:"n"`
				Script  string `json:"script"`
			} `json:"prev_out"`
			Script string `json:"script"`
		} `json:"inputs"`
		Time      int    `json:"time"`
		TtxIndex  int    `json:"ttx_index"`
		VinSz     int    `json:"vin_sz"`
		Hash      string `json:"hash"`
		VoutSz    int    `json:"vout_sz"`
		RelayedBy string `json:"relayed_by"`
		Out       []struct {
			Spent   bool   `json:"spent"`
			TxIndex int    `json:"tx_index"`
			Type    int    `json:"type"`
			Addr    string `json:"addr"`
			Value   int    `json:"value"`
			N       int    `json:"n"`
			Script  string `json:"script"`
		} `json:"out"`
	} `json:"x"`
}



func main() {

	for {
		c := evtwebsocket.Conn{

			// When connection is established
			OnConnected: func(w *evtwebsocket.Conn) {
				log.Println("Connected")
			},

			// When a message arrives
			OnMessage: func(msg []byte, w *evtwebsocket.Conn) {
				//log.Printf("OnMessage: %s\n", msg)
				//fmt.Println(msg)
				fmt.Printf("MESSAGE: %s\n", msg)
				var transact Transaction
				//data := []byte(msg)
				err := json.Unmarshal(msg, &transact)
				if err == nil {
					//fmt.Printf("%s", msg)
					fmt.Printf("INPUTS:\n")
					for k, v := range transact.X.Inputs {
						//fmt.Println(parsed["inputs"])
						fmt.Println(k, v)
					}
					fmt.Println("")
				}
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
	}
}
