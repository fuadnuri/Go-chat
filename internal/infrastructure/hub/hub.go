package hub

import (
	"fmt"

	"github.com/fuadnuri/Go-chat/ltypes"
)

type Hub struct {
	Clients     map[*Client]bool
	Subscribe   chan *Client
	Unsubscribe chan *Client
	DM          chan ltypes.DirectMessage
}

func NewHub() *Hub {
	return &Hub{
		Clients:     make(map[*Client]bool),
		Subscribe:   make(chan *Client),
		Unsubscribe: make(chan *Client),
		DM:          make(chan ltypes.DirectMessage),
	}
}

func (h *Hub) Run() {
	defer func() {
		close(h.Subscribe)
		close(h.Unsubscribe)
		close(h.DM)
		//wg.done()
	}()
	for {
		select {
		case client, ok := <-h.Subscribe:
			if ok {
				fmt.Println(client.User.Username)
				if !h.Clients[client] {
					h.Clients[client] = true

				} else {
					continue
				}

			}
		case client, ok := <-h.Unsubscribe:
			if ok {
				if h.Clients[client] {
					client.Conn.Close()
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		case message := <-h.DM:
			for client := range h.Clients {
				if message.SenderUsername == client.User.Username {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(h.Clients, client)
					}
				}

			}
		}
	}
}
