# WebSocket Server with Fiber

This project demonstrates how to create a WebSocket server using the Fiber framework in Go
The server maintains all connections in a variable and broadcasts messages to all connections except the sender.

## Prerequisites

- Go 1.22 or later
- Fiber v2

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/websocket-server-fiber.git
   cd websocket-server-fiber
   ```

2. Install the required packages:
   ```bash
    go get github.com/gofiber/fiber/v2
    go get github.com/gofiber/websocket/v2
    ```

3. Run the application:
    ```bash
    go run main.go
    ```