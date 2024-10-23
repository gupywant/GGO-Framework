// config/db.go
package config

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
	"github.com/joho/godotenv"
)

// DBs stores multiple DB connections, keyed by the database name.
var DBs = make(map[string]*gorm.DB)
var GGODB *gorm.DB

// ConnectDB connects to the specified database and stores it in the DBs map.
func ConnectDB() {
	
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbName := os.Getenv("DB_NAME") // Fetch the DB name dynamically

    // Fetch other environment variables
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    port := os.Getenv("DB_PORT")
	defaultConnectionName := os.Getenv("DB_CONNECTION")

    // Create the connection string for PostgreSQL
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

    // Connect to the database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database %s. Error: %s", dbName, err)
    }

    // Store the connection in the DBs map, keyed by the database name
    DBs[defaultConnectionName] = db
	
	if _, exists := DBs[defaultConnectionName]; exists {
        GGODB = DBs[defaultConnectionName]
    } else {
        log.Fatalf("Default DB connection %s does not exist in DBs map", defaultConnectionName)
    }

    fmt.Printf("Connected to database: %s\n", dbName)
}

// GetDB returns the *gorm.DB connection for a given database name
func GetDB(connectionName string) *gorm.DB {
	if connectionName == "" {
		return GGODB
	}
    return DBs[connectionName]
}
