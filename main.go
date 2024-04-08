package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define a route handler for the homepage
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, please enter your name: <form method='post' action='/save'><input type='text' name='name'><input type='submit' value='Submit'></form>")
	})

	// Define a route handler to save the name
	app.Post("/save", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		// Save the name to the database or perform any other necessary action
		log.Printf("Name: %s", name)
		// Return a response indicating the name has been saved
		return c.SendString("Thank you, " + name + ", your name has been saved!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
