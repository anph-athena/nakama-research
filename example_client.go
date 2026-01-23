package main

import (
	"fmt"
)

// This is an example client that demonstrates how to use the level-based matchmaking
// In a real scenario, this would be a separate client application

func main() {
	fmt.Println("Example matchmaking client")
	fmt.Println("To use this client, you need to:")
	fmt.Println("1. Run 'docker-compose up -d' to start Nakama server")
	fmt.Println("2. Wait for the server to be ready")
	fmt.Println("3. Use the Nakama client SDK to connect and add matchmaking tickets")
	fmt.Println("")
	fmt.Println("Example matchmaking ticket with level:")
	fmt.Println(`
	ticket, err := client.AddMatchmaker(ctx, 2, 8, "*", map[string]string{
		"level": "10",
	}, map[string]float64{
		"level": 10.0, // Player's level for matchmaking
	})
	`)
	fmt.Println("")
	fmt.Println("The matchmaker will group players with similar levels together")
	fmt.Println("based on the level_range configured in local.yml")
}

// ExampleMatchmaking demonstrates how to add a matchmaking ticket with level
func ExampleMatchmaking() {
	// This is pseudocode to show the concept
	// In a real implementation, you would use the official Nakama Go client SDK

	// Create a matchmaking ticket with the player's level
	playerLevel := 10.0

	// The matchmaker will use the "level" property to match players
	// Players with similar levels will be matched together
	stringProps := map[string]string{
		"region": "us-east",
	}

	numericProps := map[string]float64{
		"level": playerLevel, // This is the key property for matchmaking
	}

	fmt.Printf("Adding matchmaker ticket for player with level %.0f\n", playerLevel)
	fmt.Println("String properties:", stringProps)
	fmt.Println("Numeric properties:", numericProps)

	// In a real client, you would call something like:
	// ticket, err := client.AddMatchmaker(ctx, 2, 8, "*", stringProps, numericProps)

	fmt.Println("Matchmaking ticket added (example)")

	// The server's matchmakerMatched function will be called when players are matched
}

// Example of creating authenticated sessions
func ExampleAuthentication() {
	fmt.Println("\nAuthentication Example:")
	fmt.Println("To authenticate with Nakama, use device ID or custom authentication:")
	fmt.Println(`
	// Using device authentication
	session, err := client.AuthenticateDevice(ctx, deviceID, true, "Player1")
	
	// Using custom authentication  
	session, err := client.AuthenticateCustom(ctx, customID, true, "Player1")
	`)
}

// Example match configuration
func ExampleMatchConfig() {
	fmt.Println("\nMatch Configuration:")
	fmt.Println("The matchmaker uses the following rules from local.yml:")
	fmt.Println("- min_party_size: 2 (minimum players per match)")
	fmt.Println("- max_party_size: 8 (maximum players per match)")
	fmt.Println("- interval_sec: 5 (matchmaking runs every 5 seconds)")
	fmt.Println("- level_range: 5 (players within 5 levels can be matched)")
	fmt.Println("")
	fmt.Println("The matchmaker will group players whose levels are within the configured range.")
}
