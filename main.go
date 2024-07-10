package main

import (
	"log"
	"server/pool"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := pool.Init()
	if err != nil {
		log.Fatalln(err)
	}
	app := fiber.New()
	app.Use("/ws", pool.Upgrade)
	app.Get("/ws/feed", websocket.New(pool.Handler))
	err = app.Listen(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
