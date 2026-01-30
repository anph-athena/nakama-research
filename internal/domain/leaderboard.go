package domain

import "context"

// LeaderboardEntry represents a player's entry in the leaderboard
type LeaderboardEntry struct {
	Rank        int64
	PlayerID    string
	Username    string
	DisplayName string
	Score       int64
	UpdatedAt   int64
}

// Leaderboard represents the leaderboard aggregate
type Leaderboard struct {
	ID        string
	Name      string
	Entries   []*LeaderboardEntry
	UpdatedAt int64
}

// LeaderboardRepository defines the interface for leaderboard data access
type LeaderboardRepository interface {
	GetLeaderboard(ctx context.Context, limit, offset int) ([]*LeaderboardEntry, error)
	GetPlayerRank(ctx context.Context, playerID string) (int64, error)
	UpdateScore(ctx context.Context, playerID string, score int64) error
	GetTopN(ctx context.Context, n int) ([]*LeaderboardEntry, error)
	GetAroundRank(ctx context.Context, playerID string, range_ int) ([]*LeaderboardEntry, error)
}
