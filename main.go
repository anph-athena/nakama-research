package main

import (
	"context"
	"database/sql"
	"nakama-research/internal/platform"

	"github.com/heroiclabs/nakama-common/runtime"
)

// InitModule initializes the Nakama module
// This is the entry point called by Nakama when loading the plugin
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	// Create module initializer with all dependencies
	module, err := platform.NewModuleInitializer(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error("Failed to initialize Nakama module: %v", err)
		return err
	}

	// Register all RPC handlers
	if err := module.RegisterRPCs(ctx); err != nil {
		logger.Error("Failed to register RPCs: %v", err)
		return err
	}

	logger.Info("Nakama module loaded successfully")
	return nil
}
