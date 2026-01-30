package redis

import (
	"context"
	"nakama-research/internal/domain"
)

type RedisLeaderboardRepository struct{}

func NewRedisLeaderboardRepository() *RedisLeaderboardRepository {
	return &RedisLeaderboardRepository{}
}

func (r *RedisLeaderboardRepository) GetLeaderboard(ctx context.Context, limit, offset int) ([]*domain.LeaderboardEntry, error) {
	return []*domain.LeaderboardEntry{}, nil
}

func (r *RedisLeaderboardRepository) GetPlayerRank(ctx context.Context, playerID string) (int64, error) {
	return 0, nil
}

func (r *RedisLeaderboardRepository) UpdateScore(ctx context.Context, playerID string, score int64) error {
	return nil
}

func (r *RedisLeaderboardRepository) GetTopN(ctx context.Context, n int) ([]*domain.LeaderboardEntry, error) {
	return []*domain.LeaderboardEntry{}, nil
}

func (r *RedisLeaderboardRepository) GetAroundRank(ctx context.Context, playerID string, range_ int) ([]*domain.LeaderboardEntry, error) {
	return []*domain.LeaderboardEntry{}, nil
}
