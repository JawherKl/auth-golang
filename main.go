// PATH: auth-golang/main.go

package main

import (
    "auth-golang/models"
    "auth-golang/routes"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Create a new gin instance
    r := gin.Default()

    // Load .env file and Create a new connection to the database
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize DB
    models.InitDB()

    // Load the routes
    routes.AuthRoutes(r)

    // Run the server
    r.Run(":8080")
}
