package pool

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Init() (err error) {
	conns := make(map[*websocket.Conn]bool)
	connPool = connection{
		connection: conns,
		mu:         &sync.Mutex{},
	}
	return
}

func (conn *connection) AddConnection(connection *websocket.Conn) {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	conn.connection[connection] = true
}

func (conn *connection) DeleteConnection(connection *websocket.Conn) {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	delete(conn.connection, connection)
}

func Upgrade(c *fiber.Ctx) (err error) {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (conn *connection) BroadCast(sender *websocket.Conn, message []byte) (err error) {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	for conn := range conn.connection {
		if conn != sender {
			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				delete(connPool.connection, conn)
				conn.Close()
			}
		}
	}
	return
}

func Handler(conn *websocket.Conn) {
	connPool.AddConnection(conn)
	defer func() {
		connPool.DeleteConnection(conn)
	}()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		err = connPool.BroadCast(conn, msg)
		if err != nil {
			break
		}
	}
	return
}
