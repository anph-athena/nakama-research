package domain

import "context"

// Player represents the player entity
type Player struct {
	ID          string
	Username    string
	CustomID    string
	Email       string
	DisplayName string
	Score       int64
	Level       int32
	LastActive  int64
	CreatedAt   int64
	UpdatedAt   int64
}

// PlayerRepository defines the interface for player data access
type PlayerRepository interface {
	GetByID(ctx context.Context, id string) (*Player, error)
	GetByCustomID(ctx context.Context, customID string) (*Player, error)
	GetByUsername(ctx context.Context, username string) (*Player, error)
	Create(ctx context.Context, player *Player) (*Player, error)
	Update(ctx context.Context, player *Player) (*Player, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*Player, error)
	CountAll(ctx context.Context) (int64, error)
}
