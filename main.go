package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Todo struct {
	Id   int8
	Name string
	Done bool
}

var todos []Todo = []Todo{
	{Name: "Create simple todo app", Done: false, Id: 0},
	{Name: "Learn basics of Golang", Done: true, Id: 1},
}

func main() {
	// Set up template engine
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Todos": todos,
		})
	})

	// Listen on port 3000
	log.Fatal(app.Listen(":3000"))
}
