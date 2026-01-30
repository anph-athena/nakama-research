package platform

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

// ModuleInitializer holds all dependencies for the Nakama module
type ModuleInitializer struct {
	DB          *sql.DB
	Logger      runtime.Logger
	Nakama      runtime.NakamaModule
	Initializer runtime.Initializer
}

// NewModuleInitializer creates and initializes the module with all dependencies
func NewModuleInitializer(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) (*ModuleInitializer, error) {
	logger.Info("Initializing Nakama module...")

	// Initialize database tables
	if err := InitializeTables(db); err != nil {
		logger.Error("Failed to initialize database tables: %v", err)
		return nil, fmt.Errorf("failed to initialize database tables: %w", err)
	}

	m := &ModuleInitializer{
		DB:          db,
		Logger:      logger,
		Nakama:      nk,
		Initializer: initializer,
	}

	logger.Info("Nakama module initialized successfully")
	return m, nil
}

// RegisterRPCs registers all RPC handlers
func (m *ModuleInitializer) RegisterRPCs(ctx context.Context) error {
	m.Logger.Info("Registering RPC handlers...")

	// Example: Register player RPC handlers
	// You can add RPC registrations here as needed

	return nil
}

// Close closes all resources
func (m *ModuleInitializer) Close() error {
	if m.DB != nil {
		return m.DB.Close()
	}
	return nil
}
