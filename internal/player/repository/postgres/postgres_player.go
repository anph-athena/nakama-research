package postgres

import (
	"context"
	"database/sql"
	"nakama-research/internal/domain"
)

type PostgresPlayerRepository struct {
	db *sql.DB
}

func NewPostgresPlayerRepository(db *sql.DB) *PostgresPlayerRepository {
	return &PostgresPlayerRepository{db: db}
}

func (r *PostgresPlayerRepository) GetByID(ctx context.Context, id string) (*domain.Player, error) {
	return nil, nil
}

func (r *PostgresPlayerRepository) GetByCustomID(ctx context.Context, customID string) (*domain.Player, error) {
	return nil, nil
}

func (r *PostgresPlayerRepository) GetByUsername(ctx context.Context, username string) (*domain.Player, error) {
	return nil, nil
}

func (r *PostgresPlayerRepository) Create(ctx context.Context, player *domain.Player) (*domain.Player, error) {
	return player, nil
}

func (r *PostgresPlayerRepository) Update(ctx context.Context, player *domain.Player) (*domain.Player, error) {
	return player, nil
}

func (r *PostgresPlayerRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *PostgresPlayerRepository) List(ctx context.Context, limit, offset int) ([]*domain.Player, error) {
	return []*domain.Player{}, nil
}

func (r *PostgresPlayerRepository) CountAll(ctx context.Context) (int64, error) {
	return 0, nil
}
