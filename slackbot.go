package slackbot

import "golang.org/x/net/websocket"

type RedditProcessor interface {
	GetRedditPosts()
}

type SlackProcessor interface {
	SlackStart() (string, string, error)
	SlackConnect() (*websocket.Conn, string)
	GetMessage(ws *websocket.Conn) (Message, error)
	PostMessage(ws *websocket.Conn, m Message) error
}

// These are the messages read off and written into the websocket. Since this
// struct serves as both read and write, we include the "Id" field which is
// required only for writing.

type Message struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}
