package http

import (
	"nakama-research/internal/leaderboard/usecase"
	"net/http"
)

type LeaderboardHandler struct {
	usecase *usecase.LeaderboardUsecase
}

func NewLeaderboardHandler(uc *usecase.LeaderboardUsecase) *LeaderboardHandler {
	return &LeaderboardHandler{usecase: uc}
}

func (h *LeaderboardHandler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
}
