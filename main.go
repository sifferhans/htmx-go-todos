package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
)

type Todo struct {
	Id   uuid.UUID
	Name string
	Done bool
}

var todos []Todo = []Todo{
	{Name: "Create simple todo app", Done: false, Id: uuid.New()},
	{Name: "Learn basics of Golang", Done: true, Id: uuid.New()},
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

	app.Post("/add", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		todo := Todo{Name: name, Done: false, Id: uuid.New()}
		todos = append(todos, todo)

		return c.Render("todo", fiber.Map{
			"Todo": todo,
		})
	})

	app.Post("/delete", func(c *fiber.Ctx) error {
		return c.SendString("")
	})

	// Listen on port 3000
	log.Fatal(app.Listen(":3000"))
}
