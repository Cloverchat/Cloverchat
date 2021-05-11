package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	Address string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *Server) Run() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		conn.WriteJSON(ConstructMessage(&ResponseContent{
			Data: "Hello!",
		}))
		go func() {
			for {
				// Read message from browser
				resp := ClientResponse{}
				err := conn.ReadJSON(&resp)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(resp.Code == nil)

				// Print the message to the console
				fmt.Printf("%+v\n", resp)
				//fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

				// Write message back to browser
				//if err = conn.WriteMessage(msgType, msg); err != nil {
				//	return
				//}
			}
		}()
	})

	http.ListenAndServe(s.Address, nil)
}
