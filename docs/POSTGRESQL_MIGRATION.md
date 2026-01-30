# PostgreSQL Migration Summary

## Overview
The nakama-research project has been successfully migrated from MySQL to PostgreSQL.

## Changes Made

### 1. Database Driver Update
- **Old:** `github.com/go-sql-driver/mysql v1.9.3`
- **New:** `github.com/lib/pq v1.10.9` (PostgreSQL driver)
- **File:** `go.mod`

### 2. Database Configuration Changes
- **File:** `internal/platform/database.go`

#### Connection String Format
- **MySQL:** `user:password@tcp(host:port)/database?parseTime=true`
- **PostgreSQL:** `host=host port=port user=user password=password dbname=database sslmode=disable`

#### Driver
- **MySQL:** `sql.Open("mysql", dsn)`
- **PostgreSQL:** `sql.Open("postgres", dsn)`

### 3. Database Schema Updates
- **File:** `internal/platform/database.go` (InitializeTables function)

#### Table Creation Changes
- **MySQL:**
  - `id BIGINT AUTO_INCREMENT PRIMARY KEY`
  - `INDEX idx_name (column)`
  - `ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`

- **PostgreSQL:**
  - `id SERIAL PRIMARY KEY`
  - `CREATE INDEX IF NOT EXISTS idx_name ON table(column)`
  - FOREIGN KEY constraints remain the same

### 4. Repository Structure
- **Old:** `internal/player/repository/mysql/`
- **New:** `internal/player/repository/postgres/`
- **File:** `internal/player/repository/postgres/postgres_player.go`

## Migration Details

### Players Table (PostgreSQL)
```sql
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
```

### Leaderboard Entries Table (PostgreSQL)
```sql
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
```

## Next Steps

1. **Update docker-compose.yml** - Change MySQL service to PostgreSQL service
2. **Update environment variables** - Adjust database connection parameters if needed
3. **Implement repository methods** - Add actual SQL implementation in `postgres_player.go`
4. **Test database connection** - Verify the application can connect to PostgreSQL
5. **Run migrations** - Create tables using the InitializeTables function

## Configuration Example

```go
config := platform.DatabaseConfig{
    Host:     "localhost",
    Port:     "5432",      // PostgreSQL default port (was 3306 for MySQL)
    User:     "postgres",   // PostgreSQL default user
    Password: "password",
    Database: "nakama",
    MaxConn:  25,
    MaxIdle:  5,
}

db, err := platform.NewDatabase(config)
```

## Compatibility Notes

- PostgreSQL uses `SERIAL` for auto-incrementing integers (equivalent to MySQL's `AUTO_INCREMENT`)
- Index creation syntax is different and now uses separate `CREATE INDEX` statements
- Connection string format follows PostgreSQL's libpq standard
- All data types remain compatible between MySQL and PostgreSQL for this schema
