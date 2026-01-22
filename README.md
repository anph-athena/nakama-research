# nakama-research

This repository contains a Nakama server implementation with a level-based matchmaking function written in Go.

## Overview

Nakama is an open-source server for building social and real-time games. This project demonstrates:
- Installing and configuring Nakama via Go
- Implementing a matchmaking function that uses `player_profile.level` as the matchmaking rule
- Running Nakama with Docker Compose

## Prerequisites

- Go 1.19 or higher
- Docker and Docker Compose
- Basic understanding of Nakama concepts

## Installation

### 1. Clone the Repository

```bash
git clone <repository-url>
cd nakama-research
```

### 2. Install Go Dependencies

```bash
go mod download
```

### 3. Start Nakama Server

Start the Nakama server and CockroachDB database using Docker Compose:

```bash
docker-compose up -d
```

Wait for the services to be healthy. You can check the status with:

```bash
docker-compose ps
```

### 4. Verify Installation

The Nakama server should be running on:
- Console: http://localhost:7351 (admin/password)
- Client API: http://localhost:7349
- gRPC: localhost:7350

## Matchmaking Function

### How It Works

The matchmaking function (`matchmakerMatched` in `main.go`) uses the player's level as the primary matchmaking criterion:

1. **Player Properties**: When players join matchmaking, they include their level in the numeric properties
2. **Level-Based Matching**: The matchmaker groups players with similar levels together
3. **Level Range**: Configured in `local.yml` (default: 5 levels difference)
4. **Match Creation**: When suitable players are found, the function validates the level range and creates a match

### Key Features

- **Level Validation**: Checks that all matched players are within an acceptable level range
- **Logging**: Detailed logging of player levels and match creation
- **Flexible**: Can be extended to include additional matchmaking criteria

### Configuration

The matchmaking behavior is configured in `local.yml`:

```yaml
matchmaker:
  max_tickets: 1000
  interval_sec: 5          # Matchmaking runs every 5 seconds
  max_intervals: 20
  min_party_size: 2        # Minimum 2 players per match
  max_party_size: 8        # Maximum 8 players per match
  level_range: 5           # Players within 5 levels can match
```

## Using the Matchmaker

### Client Implementation

To use the matchmaker from a game client, you need to:

1. **Authenticate**: Create a session with the Nakama server
2. **Add Matchmaker Ticket**: Submit a matchmaking request with your level

Example (pseudocode):

```go
// Authenticate
session, err := client.AuthenticateDevice(ctx, deviceID, true, "PlayerName")

// Add matchmaker ticket with level
ticket, err := client.AddMatchmaker(
    ctx,
    2,                      // min players
    8,                      // max players
    "*",                    // query (matches all)
    map[string]string{      // string properties
        "region": "us-east",
    },
    map[string]float64{     // numeric properties
        "level": 10.0,      // Player level for matchmaking
    },
)

// Wait for match to be found
// The matchmakerMatched function will be called on the server
```

### Testing Matchmaking

You can test the matchmaking using the Nakama Console or client SDKs:

1. Open the Nakama Console: http://localhost:7351
2. Login with username `admin` and password `password`
3. Create test users and add matchmaker tickets with different levels
4. Monitor the server logs to see matches being created

View logs:
```bash
docker-compose logs -f nakama
```

## Project Structure

```
nakama-research/
├── main.go                 # Nakama server module with matchmaking function
├── example_client.go       # Example client code demonstrating usage
├── local.yml              # Nakama server configuration
├── docker-compose.yml     # Docker setup for Nakama and CockroachDB
├── go.mod                 # Go module dependencies
├── go.sum                 # Go module checksums
└── README.md              # This file
```

## Architecture

### Components

1. **CockroachDB**: Distributed SQL database for Nakama
2. **Nakama Server**: Game server handling authentication, matchmaking, and more
3. **Go Runtime Module**: Custom server-side logic (matchmaking function)

### Matchmaking Flow

```
1. Player connects to Nakama
2. Player authenticates
3. Player adds matchmaker ticket with level property
4. Matchmaker groups players with similar levels (every 5 seconds)
5. matchmakerMatched function is called
6. Function validates level range and creates match
7. Players are notified of the match
```

## Development

### Modifying the Matchmaking Logic

Edit `main.go` to customize the matchmaking behavior:

```go
func matchmakerMatched(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {
    // Your custom matchmaking logic here
    // Access player levels via entry.GetProperties()["level"]
    return "", nil
}
```

After making changes:

```bash
# Rebuild and restart
docker-compose down
docker-compose up -d
```

### Adding More Matchmaking Criteria

You can extend the matchmaking to include additional properties:

```go
numericProps := map[string]float64{
    "level": playerLevel,
    "skill": playerSkill,
    "rank": playerRank,
}
```

Then update the `matchmakerMatched` function to handle these properties.

## Troubleshooting

### Server Won't Start

Check logs:
```bash
docker-compose logs nakama
docker-compose logs cockroachdb
```

### Matchmaking Not Working

1. Verify players are adding tickets with the "level" property
2. Check that min_party_size is met (at least 2 players)
3. Wait for the matchmaking interval (5 seconds by default)
4. Review logs for matchmaker activity

### Database Connection Issues

Ensure CockroachDB is healthy before Nakama starts:
```bash
docker-compose ps
```

## References

- [Nakama Documentation](https://heroiclabs.com/docs/)
- [Nakama Server Setup](https://heroiclabs.com/docs/nakama/getting-started/install/macos/)
- [Matchmaker Documentation](https://heroiclabs.com/docs/nakama/tutorials/unity/pirate-panic/matchmaking/#requesting-a-match)
- [Go Runtime Documentation](https://heroiclabs.com/docs/nakama/server-framework/go-runtime/)

## License

This project is for research and educational purposes.
