package cli

import (
	"fmt"

	"github.com/randyazharalman/blog_aggregator/internal/config"
	"github.com/spf13/cobra"
)

type State struct {
	Config *config.Config
}

var state *State

var rootCmd = &cobra.Command{
	Use: "gator",
	Short: "A blog aggregator CLI tool",
	Long: `Gator is a CLI tool for managing and aggregating blog feeds.

It allows you to:
- Login with a username
- Add RSS/Atom feeds to follow
- Browse and read blog posts
- Manage your feed subscriptions`,
	PersistentPreRunE: initializeState,
}

func Execute() error {
	return rootCmd.Execute()
}

func initializeState(cmd *cobra.Command, args []string) error {
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("failed to read configuration: %w", err)
	}

	state = &State{
		Config: &cfg,
	}

	return nil
}


func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(resetCmd)
}