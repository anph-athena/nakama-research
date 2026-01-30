package domain

import "context"

// AuthCredentials represents authentication credentials
type AuthCredentials struct {
	UserID    string
	CustomID  string
	Token     string
	ExpiresAt int64
}

// AuthService defines the interface for authentication operations
type AuthService interface {
	ValidateToken(ctx context.Context, token string) (*AuthCredentials, error)
	LoginByCustomID(ctx context.Context, customID string) (*AuthCredentials, error)
	RefreshToken(ctx context.Context, token string) (*AuthCredentials, error)
}
