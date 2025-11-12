package main
import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    fmt.Println("Starting server on :8080") 
    if err := app.Listen(":8080"); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

