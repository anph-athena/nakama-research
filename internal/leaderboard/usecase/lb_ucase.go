package usecase

import (
	"context"
	"nakama-research/internal/domain"
)

type LeaderboardUsecase struct {
	repo domain.LeaderboardRepository
}

func NewLeaderboardUsecase(repo domain.LeaderboardRepository) *LeaderboardUsecase {
	return &LeaderboardUsecase{repo: repo}
}

func (uc *LeaderboardUsecase) GetLeaderboard(ctx context.Context, limit, offset int) ([]*domain.LeaderboardEntry, error) {
	return uc.repo.GetLeaderboard(ctx, limit, offset)
}
