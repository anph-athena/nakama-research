# Testing Guide for Level-Based Matchmaking

This guide explains how to test the level-based matchmaking functionality in the Nakama server.

## Quick Start

1. **Start the server**:
   ```bash
   ./setup.sh
   # Or manually:
   docker-compose up -d
   ```

2. **Access the Console**: Open http://localhost:7351 in your browser
   - Username: `admin`
   - Password: `password`

## Testing Matchmaking

### Option 1: Using Nakama Console (Web UI)

1. Navigate to the Nakama Console at http://localhost:7351
2. Log in with admin credentials
3. Go to "Users" section and create test users
4. Use the API Explorer to add matchmaker tickets for each user

### Option 2: Using cURL

Create matchmaker tickets with different levels:

```bash
# First, authenticate to get a session token
# Note: Replace with actual authentication endpoint

# Example matchmaker ticket (requires valid session token)
curl -X POST 'http://localhost:7349/v2/matchmaker' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer YOUR_SESSION_TOKEN' \
  -d '{
    "min_count": 2,
    "max_count": 8,
    "query": "*",
    "string_properties": {},
    "numeric_properties": {
      "level": 10
    }
  }'
```

### Option 3: Using Nakama Client SDK

Install a Nakama client SDK for your preferred language:

**JavaScript/TypeScript:**
```bash
npm install @heroiclabs/nakama-js
```

**Go:**
```bash
go get github.com/heroiclabs/nakama-common@latest
```

**Example client code (JavaScript):**
```javascript
const { Client } = require("@heroiclabs/nakama-js");

const client = new Client("defaultkey", "localhost", "7349");
client.ssl = false;

async function testMatchmaking() {
  // Authenticate
  const session = await client.authenticateDevice("device-id-1", true);
  
  // Add matchmaker ticket with level
  const socket = client.createSocket();
  await socket.connect(session);
  
  const ticket = await socket.addMatchmaker({
    minCount: 2,
    maxCount: 8,
    query: "*",
    stringProperties: {},
    numericProperties: {
      level: 10  // Player level
    }
  });
  
  console.log("Matchmaker ticket:", ticket);
  
  // Listen for match found
  socket.onmatchmakermatched = (matched) => {
    console.log("Match found:", matched);
  };
}

testMatchmaking();
```

## Test Scenarios

### Scenario 1: Matching Players with Similar Levels

**Goal**: Verify that players with similar levels are matched together.

**Steps**:
1. Create Player A with level 10
2. Create Player B with level 12
3. Add matchmaker tickets for both
4. Wait ~5 seconds (matchmaking interval)
5. Check logs: Both players should be matched together

**Expected Result**: Players A and B are matched (within level_range: 5)

### Scenario 2: Players with Different Levels

**Goal**: Verify that players with very different levels are NOT matched.

**Steps**:
1. Create Player A with level 10
2. Create Player B with level 20
3. Add matchmaker tickets for both
4. Wait for matchmaking

**Expected Result**: Players wait longer or need more players within their level range

### Scenario 3: Multiple Players at Same Level

**Goal**: Verify efficient matching of similar-level players.

**Steps**:
1. Create 4 players all at level 15
2. Add matchmaker tickets for all
3. Wait ~5 seconds

**Expected Result**: All 4 players are matched into a single game

## Monitoring

### View Server Logs

```bash
docker-compose logs -f nakama
```

Look for these log messages:
```
Nakama server module initialized successfully with level-based matchmaking
Matchmaker matched called with N entries
Player 1 matched: <user_id>
Player <user_id> level: <level>
Match created with level range: <range> (min: <min>, max: <max>)
```

### Check CockroachDB

```bash
docker-compose exec cockroachdb ./cockroach sql --insecure --database=nakama

# View matchmaker tickets
SELECT * FROM matchmaker_tickets;
```

## Troubleshooting

### Matchmaking Not Working

1. **Check minimum players**: Ensure at least `min_party_size` (2) players have added tickets
2. **Wait for interval**: Matchmaking runs every `interval_sec` (5) seconds
3. **Verify levels**: Check that player levels are within `level_range` (5)
4. **Check logs**: Look for errors in `docker-compose logs nakama`

### No Matches Created

- Verify the `matchmakerMatched` function is being called
- Check that players are adding tickets with the "level" property
- Ensure numeric properties are being passed correctly (not as strings)

### Server Won't Start

1. Stop all services: `docker-compose down`
2. Remove volumes: `docker-compose down -v`
3. Start fresh: `docker-compose up -d`

## Configuration

Modify matchmaking behavior in `local.yml`:

```yaml
matchmaker:
  max_tickets: 1000        # Maximum concurrent tickets
  interval_sec: 5          # How often matchmaking runs
  max_intervals: 20        # Max wait time (intervals * interval_sec)
  min_party_size: 2        # Minimum players per match
  max_party_size: 8        # Maximum players per match
  level_range: 5           # Allowed level difference
```

After changes, restart:
```bash
docker-compose restart nakama
```

## Advanced Testing

### Load Testing

Create a script to simulate many players:

```bash
for i in {1..10}; do
  # Create player with random level between 1-20
  level=$((1 + RANDOM % 20))
  echo "Creating player $i with level $level"
  # Add matchmaker ticket (requires valid auth)
done
```

### Custom Match Properties

Test additional properties:

```javascript
numericProperties: {
  level: 10,
  skill_rating: 1500,
  ping: 50
}
```

Update `main.go` to use these properties in matchmaking logic.

## Next Steps

1. Implement custom match logic based on match results
2. Add more sophisticated matchmaking rules (skill rating, region, etc.)
3. Create a real-time match handler
4. Implement leaderboards based on match outcomes

## Resources

- [Nakama Matchmaker Docs](https://heroiclabs.com/docs/nakama/concepts/matchmaker/)
- [Client SDK Reference](https://heroiclabs.com/docs/nakama/client-libraries/)
- [Server Runtime Docs](https://heroiclabs.com/docs/nakama/server-framework/go-runtime/)
