// File: /server/server.go
package server

import (
	"log"

	"github.com/danfestudio/youthcongress.org.np/database"
	"github.com/gofiber/fiber/v2"
)

// StartServer initializes and starts the Fiber server
func StartServer() {	
	// Initialize database connection
	database.Connection()
	defer database.Close()

	// Ensure the `members` table exists
	database.CreateMemberTable(database.DB)

	// Create a new Fiber app
	app := fiber.New()

	// Serve static files (CSS, JS, Images)
	app.Static("/", "./public")

	// Register routes
	urlRoutes(app)

	// Register application routes
	RegisterRoutes(app)

	// Start the server
	port := ":8011"
	log.Printf("Starting server on http://localhost%s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
