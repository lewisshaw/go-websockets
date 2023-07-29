package sockets

import (
	"lewisshaw/go-ws/models"

	"github.com/gorilla/websocket"
)

type Client interface {
	getConn() *websocket.Conn
	getRace() models.Race
}

type client struct {
	conn *websocket.Conn
	race models.Race
}

func NewClient(conn *websocket.Conn, race models.Race) *client {
	return &client{
		conn: conn,
		race: race,
	}
}

func (c *client) getConn() *websocket.Conn {
	return c.conn
}

func (c *client) getRace() models.Race {
	return c.race
}
