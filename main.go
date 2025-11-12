package main
import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()


	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "ok", "message": "Service is healthy"})
	})

    fmt.Println("Starting server on :8080") 
    if err := app.Listen(":8080"); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

