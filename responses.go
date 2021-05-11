// Here we define server and client response structs.
package main

import "time"
import "strconv"

// content field of responses
type ResponseContent struct {
	Data string `json:"data"`
	Files []string `json:"files,omitempty"`
}

// The server response struct
// What we should send, as defined by the procotol
type ServerResponse struct {
	Ms string `json:"ms"`
	Code string `json:"code"`
	Method string `json:"method"`
	Content ResponseContent `json:"content"`
}

// Client response struct
// This borrows a lot from a server's response
// Fields we should expect from the client
type ClientResponse struct {
	SessionID string `json:"sessionID,omitempty"`
	// Code is the only required field by the spec
	// To check if it was provided, we make it a pointer to a string
	// If the pointer is nil, no `code` sent and if not nil `code` was sent
	Code *string `json:"code"`
	Method string `json:"method,omitempty"`
	Destinations []string `json:"destinations,omitempty"`
	Key string `json:"key,omitempty"`
	Content ResponseContent `json:"content,omitempty"`
}

// Utility function which constructs a ServerResponse, where Code is MESSAGE
// which is a normal message in the TermTalk protocol
func ConstructMessage(content ResponseContent, mthd ...string) ServerResponse {
	method := ""
	timestamp := time.Now().Unix()

	if len(mthd) > 0 {
		method = mthd[0]
	}

	return ServerResponse {
		Ms: strconv.FormatInt(timestamp, 10),
		Code: "MESSAGE",
		Method: method,
		Content: content,
	}
}
