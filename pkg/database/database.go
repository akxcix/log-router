package database

import (
	"database/sql"
	"fmt"

	"github.com/akxcix/log-router/pkg/database/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Database struct {
	connection *sql.DB
}

func NewDatabase() *Database {
	connStr := "user=postgres password=WSla33snLxUx3GvF host=db.vovxnvsdamdgflmtmdhq.supabase.co port=5432 dbname=postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}

	database := &Database{
		connection: db,
	}

	return database
}

func (db *Database) Save(events []models.Event) {
	query := `INSERT INTO events (id, log) VALUES ($1, $2)`

	for _, event := range events {
		_, err := db.connection.Exec(query, uuid.New(), event.Log)
		if err != nil {
			fmt.Printf("Failed to save event to database: %v\n", err)
		}
	}
}
