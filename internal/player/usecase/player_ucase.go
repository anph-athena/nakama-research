package usecase

import (
	"context"
	"nakama-research/internal/domain"
)

type PlayerUsecase struct {
	repo domain.PlayerRepository
}

func NewPlayerUsecase(repo domain.PlayerRepository) *PlayerUsecase {
	return &PlayerUsecase{repo: repo}
}

func (uc *PlayerUsecase) GetPlayerByID(ctx context.Context, id string) (*domain.Player, error) {
	return uc.repo.GetByID(ctx, id)
}
