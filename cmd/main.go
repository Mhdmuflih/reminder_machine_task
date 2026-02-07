package main

import (
	"fmt"
	"os"
	"reminder/config"
	"reminder/routes"
	"reminder/scheduler"
	// "reminder/seed"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // load .env

	// Connect to PostgreSQL DB & migrate tables
	config.ConnectDB()

	// // Seed sample data
	// seed.SeedTasks()
	// seed.SeedReminderRules()

	// ===== Background scheduler start =====
	scheduler.StartReminderScheduler()
	// ===== Background scheduler end =====

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback
	}

	r := gin.Default()
	routes.RegisterHealthRoutes(r)
	routes.RegisterRuleRoutes(r)
	routes.RegisterTaskRoutes(r)
	routes.RegisterLogRoutes(r)

	fmt.Println("Server is running hhtp://localhost:" + port)
	r.Run(":" + port)
}
