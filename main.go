package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

// InitModule initializes the Nakama server module and registers custom functions
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	// Register the matchmaker matched function
	if err := initializer.RegisterMatchmakerMatched(matchmakerMatched); err != nil {
		return fmt.Errorf("unable to register matchmaker matched: %w", err)
	}

	logger.Info("Nakama server module initialized successfully with level-based matchmaking")
	return nil
}

// matchmakerMatched is called when the matchmaker has matched players together
// This function uses player_profile.level as the matchmaking rule
func matchmakerMatched(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {
	logger.Info("Matchmaker matched called with %d entries", len(entries))
	
	// Validate that all players are within the level range
	var minLevel, maxLevel float64
	firstEntry := true
	
	// Log the matched players and their levels
	for i, entry := range entries {
		userID := entry.GetPresence().GetUserId()
		logger.Info("Player %d matched: %s", i+1, userID)
		
		// Get the level from properties
		if props := entry.GetProperties(); props != nil {
			if levelValue, ok := props["level"]; ok {
				var level float64
				switch v := levelValue.(type) {
				case float64:
					level = v
				case int64:
					level = float64(v)
				case int:
					level = float64(v)
				default:
					logger.Warn("Unknown level type for player %s: %T", userID, levelValue)
					continue
				}
				
				logger.Info("Player %s level: %.0f", userID, level)
				
				if firstEntry {
					minLevel = level
					maxLevel = level
					firstEntry = false
				} else {
					if level < minLevel {
						minLevel = level
					}
					if level > maxLevel {
						maxLevel = level
					}
				}
			} else {
				logger.Warn("Player %s has no level property", userID)
			}
		}
	}
	
	levelRange := maxLevel - minLevel
	logger.Info("Match created with level range: %.0f (min: %.0f, max: %.0f)", levelRange, minLevel, maxLevel)
	
	// Return empty string to use default match handler
	// You can return a custom match ID here if you want to use a custom match handler
	return "", nil
}
