package cli

import (
	"fmt"
	"os"

	"github.com/randyazharalman/blog_aggregator/internal/config"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current configuration status",
	Long: `Display the current user and database configuration.

This command shows:
- Current logged in user
- Database connection string
- Configuration file location`,
	RunE: runStatus,
}


func runStatus(cmd *cobra.Command, args []string) error {
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("failed to read configuration : %w", err)
	}

	 fmt.Println("== Current Status ==")
	 fmt.Sprintf("User: %s", cfg.CurrentUserName)

	 	if cfg.CurrentUserName == "" {
		fmt.Printf("  ❌ No user logged in. Run 'gator login <username>' first.\n")
	} else {
		fmt.Printf("  ✅ Logged in as %s\n", cfg.CurrentUserName)
	}

		fmt.Printf("Database URL: %s\n", cfg.DbURL)
	
	// Show config file location
	if homeDir, err := os.UserHomeDir(); err == nil {
		fmt.Printf("Config file: %s/.gatorconfig.json\n", homeDir)
	}
	
	return nil
}