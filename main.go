package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-memdb"
)

type Todo struct {
	Name string
	Done bool
}

func main() {
	app := fiber.New()

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"todo": &memdb.TableSchema{
				Name: "todo",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	txn := db.Txn(true)

	// Seed db with a couple of todos
	todos := []*Todo{
		&Todo{Name: "Create simple todo app", Done: false},
		&Todo{Name: "Test Golang", Done: false},
	}
	for _, t := range todos {
		if err := txn.Insert("todo", t); err != nil {
			panic(err)
		}
	}
	txn.Commit()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.SendString("Todos endpoint")
	})

	log.Fatal(app.Listen(":3000"))
}
