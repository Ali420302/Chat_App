package ws

import (
	"log"
     "time"
	"github.com/gorilla/websocket"
	"github.com/google/uuid"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

type Message struct {
	ID       string `json:"messageId"`
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
	Timestamp time.Time
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			
			Content:  string(m),
			ID: uuid.NewString(),
			RoomID:   c.RoomID,
			Username: c.ID,
			Timestamp: time.Now(),
			
		}

		hub.Broadcast <- msg
	}
}
