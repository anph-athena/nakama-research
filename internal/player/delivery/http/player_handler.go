package http

import (
	"nakama-research/internal/player/usecase"
	"net/http"
)

type PlayerHandler struct {
	usecase *usecase.PlayerUsecase
}

func NewPlayerHandler(uc *usecase.PlayerUsecase) *PlayerHandler {
	return &PlayerHandler{usecase: uc}
}

func (h *PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request) {
}
