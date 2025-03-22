package db

import (
	"context"
	"log"
	"time"

	"github.com/go-pg/pg/v10"
)

// Connection is a global variable that holds the database connection.
var connection *pg.DB

// Config holds the database connection configuration.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// InitDBConnection initializes the global database connection.
func InitDBConnection(cfg Config) {
	connection = pg.Connect(&pg.Options{
		Addr:         cfg.Host + ":" + cfg.Port,
		User:         cfg.User,
		Password:     cfg.Password,
		Database:     cfg.Database,
		PoolSize:     10,
		MinIdleConns: 2,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	// Test the database connection.
	if err := testDBConnection(connection); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Connected to the database successfully.")
}

// GetConnection retrieves the global database connection.
func GetConnection() *pg.DB {
	return connection
}

// testDBConnection pings the database to ensure it's reachable.
func testDBConnection(db *pg.DB) error {
	ctx, cancel := createTimeoutContext(5 * time.Second)
	defer cancel()

	// Execute a simple query to check connectivity.
	_, err := db.ExecContext(ctx, "SELECT 1")
	return err
}

// createTimeoutContext creates a context with a timeout for database operations.
func createTimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
