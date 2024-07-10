package pool

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type connection struct {
	connection map[*websocket.Conn]bool
	mu         *sync.Mutex
}

var (
	connPool connection
)
