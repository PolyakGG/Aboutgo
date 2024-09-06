// Fiber is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go.
// Designed to ease things up for fast development with zero memory allocation and performance in mind.
// Fiber в Golang — это веб-фреймворк, вдохновленный Express.js, построенный на основе Fasthttp, самого быстрого HTTP-движка для Go.
// Он разработан для упрощения и ускорения разработки с нулевым выделением памяти и высокой производительностью.
//
// Пример использования Fiber:
package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	log.Fatal(app.Listen(":3000"))
}
