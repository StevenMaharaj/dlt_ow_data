package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

func trades() {
	msg := `{
    "type": "subscribe",
    "payload": {
        "channels": [
            {
                "name": "all_trades",
                "symbols": [
                    "put_options",
                    "call_options"
                ]
            }
        ]
    }
}`
	file, err := os.Create("trades.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	u := url.URL{Scheme: "wss", Host: "socket.delta.exchange", Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	c.WriteMessage(websocket.TextMessage, []byte(msg))
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("received: %s", message)
		file.WriteString(string(message) + "\n")
	}
}
