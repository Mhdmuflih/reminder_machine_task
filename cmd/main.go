package main

import (
	"fmt"
	"os"
	"reminder/config"
	"reminder/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // load .env

	// Connect to PostgreSQL DB & migrate tables
	config.ConnectDB()

	// Seed tasks and reminder rules
	// seed.SeedTasks()
	// seed.SeedReminderRules()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback
	}

	r := gin.Default()
	routes.RegisterHealthRoutes(r);
	routes.RegisterRuleRoutes(r);
	routes.RegisterTaskRoutes(r);
	routes.RegisterLogRoutes(r);

	fmt.Println("Server is running hhtp://localhost:" + port)
	r.Run(":" + port)
}
