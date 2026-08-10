package hub

import (
	"fmt"
	"sync"

	"github.com/fuadnuri/Go-chat/internal/domain/entities"
	"github.com/fuadnuri/Go-chat/ltypes"
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	User entities.UserEntity
	Send chan ltypes.DirectMessage
}

func (c *Client) WritePump(wg *sync.WaitGroup) {
	defer func() {
		c.Conn.Close()
		if c.Hub != nil {
			c.Hub.Unsubscribe <- c
		}
		wg.Done()
	}()
	for {
		msg := <-c.Send
		print("writing to the connection ")
		if err := c.Conn.WriteJSON(msg); err != nil {
			if err := c.Conn.PingHandler(); err != nil {
				c.Conn.Close()
			}
		}
	}
}

func (c *Client) ReadPump(wg *sync.WaitGroup) {
	defer func() {
		c.Conn.Close()
		if c.Hub != nil {
			c.Hub.Unsubscribe <- c
		}
		wg.Done()
	}()

	for {
		var msg ltypes.DirectMessage
		if err := c.Conn.ReadJSON(&msg); err != nil {
			fmt.Println("Hub is empty can't read!")
			break
		} else {
			c.Hub.DM <- msg
			//persist message
		}
	}
}
