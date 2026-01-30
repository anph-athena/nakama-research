package platform

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DatabaseConfig holds configuration for database connection
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	MaxConn  int
	MaxIdle  int
}

// NewDatabase creates a new database connection pool
func NewDatabase(config DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(config.MaxConn)
	db.SetMaxIdleConns(config.MaxIdle)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// InitializeTables creates necessary database tables if they don't exist
func InitializeTables(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS players (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		custom_id VARCHAR(255) UNIQUE,
		email VARCHAR(255),
		display_name VARCHAR(255),
		score BIGINT DEFAULT 0,
		level INT DEFAULT 1,
		last_active BIGINT,
		created_at BIGINT NOT NULL,
		updated_at BIGINT NOT NULL
	);

	CREATE INDEX IF NOT EXISTS idx_custom_id ON players(custom_id);
	CREATE INDEX IF NOT EXISTS idx_username ON players(username);
	CREATE INDEX IF NOT EXISTS idx_score ON players(score DESC);

	CREATE TABLE IF NOT EXISTS leaderboard_entries (
		id SERIAL PRIMARY KEY,
		player_id BIGINT NOT NULL,
		rank BIGINT,
		score BIGINT DEFAULT 0,
		updated_at BIGINT NOT NULL,
		FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE
	);

	CREATE INDEX IF NOT EXISTS idx_player_id ON leaderboard_entries(player_id);
	CREATE INDEX IF NOT EXISTS idx_rank ON leaderboard_entries(rank);
	CREATE INDEX IF NOT EXISTS idx_score ON leaderboard_entries(score DESC);
	`

	if _, err := db.Exec(schema); err != nil {
		return fmt.Errorf("failed to initialize tables: %w", err)
	}

	return nil
}
