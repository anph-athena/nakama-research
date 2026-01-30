# Nakama Research API Documentation

## Overview
This document describes the REST API endpoints for the Nakama Research project.

## Base URL
```
http://localhost:7350/rpc/
```

## Authentication
Currently, the API uses Nakama's built-in authentication. Include the session token in headers for authenticated requests.

## Endpoints

### Player Management

#### Get Player by ID
```
GET /player/get
Query Parameters:
  - id: string (required) - Player ID

Response:
{
  "id": "123",
  "username": "player_name",
  "custom_id": "custom_id_123",
  "email": "player@example.com",
  "display_name": "Display Name",
  "score": 1000,
  "level": 5,
  "last_active": 1704067200,
  "created_at": 1704067200,
  "updated_at": 1704067200
}
```

#### Login by Custom ID
```
POST /player/login
Request Body:
{
  "custom_id": "custom_id_123",
  "username": "player_name",
  "email": "player@example.com",
  "display_name": "Display Name"
}

Response:
{
  "id": "123",
  "username": "player_name",
  "custom_id": "custom_id_123",
  "email": "player@example.com",
  "display_name": "Display Name",
  "score": 0,
  "level": 1,
  "last_active": 1704067200,
  "created_at": 1704067200,
  "updated_at": 1704067200
}
```

#### List Players
```
GET /players
Query Parameters:
  - limit: integer (default: 10) - Number of results
  - offset: integer (default: 0) - Pagination offset

Response:
[
  {
    "id": "123",
    "username": "player_name",
    ...
  }
]
```

#### Get Player Count
```
GET /players/count

Response:
{
  "count": 150
}
```

### Leaderboard Management

#### Get Leaderboard
```
GET /leaderboard

Query Parameters:
  - limit: integer (default: 50) - Number of entries
  - offset: integer (default: 0) - Pagination offset

Response:
[
  {
    "rank": 1,
    "player_id": "123",
    "username": "top_player",
    "display_name": "Top Player",
    "score": 50000,
    "updated_at": 1704067200
  },
  ...
]
```

#### Get Player Rank
```
GET /leaderboard/rank
Query Parameters:
  - player_id: string (required) - Player ID

Response:
{
  "rank": 42
}
```

#### Get Top N Players
```
GET /leaderboard/top
Query Parameters:
  - n: integer (default: 50) - Number of top players to fetch

Response:
[
  {
    "rank": 1,
    "player_id": "123",
    ...
  }
]
```

#### Get Players Around Rank
```
GET /leaderboard/around
Query Parameters:
  - player_id: string (required) - Player ID
  - range: integer (default: 10) - Range around player's rank

Response:
[
  {
    "rank": 30,
    ...
  },
  {
    "rank": 40,
    ...
  }
]
```

#### Update Player Score
```
POST /leaderboard/score
Request Body:
{
  "player_id": "123",
  "score": 5000
}

Response:
{
  "status": "ok"
}
```

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message",
  "status": "error"
}
```

### HTTP Status Codes
- `200` - OK
- `400` - Bad Request
- `401` - Unauthorized
- `404` - Not Found
- `409` - Conflict
- `500` - Internal Server Error

## Rate Limiting
Rate limiting is managed by the Nakama server configuration.

## Future Enhancements
- [ ] OAuth2 integration
- [ ] WebSocket subscriptions
- [ ] Real-time notifications
- [ ] Advanced filtering and sorting
- [ ] Pagination cursors
