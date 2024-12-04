// Package cmd defines the CLI.
package cmd

import (
	"context"
	"fmt"
)

// Execute executes the root command.
func Execute(ctx context.Context) error {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		return fmt.Errorf("execute root command: %w", err)
	}

	return nil
}
