// PATH: auth-golang/models/index.go

package models

import (
    "fmt"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

var DB *gorm.DB

func InitDB() {
    // Get the database URL from the .env file
    dsn := os.Getenv("DATABASE_URL")
    
    // Fallback in case the environment variable is not set
    if dsn == "" {
        panic("DATABASE_URL environment variable is not set")
    }

    // Open connection to the database using the DSN
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }

    // Migrate the schema for the User model
    if err := db.AutoMigrate(&User{}); err != nil {
        panic(fmt.Sprintf("Database migration failed: %v", err))
    }

    fmt.Println("Database connection established and schema migrated")

    DB = db
}
