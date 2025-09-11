package cli

import (
	"fmt"

	"github.com/randyazharalman/blog_aggregator/internal/config"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset configuration to defaults",
	Long: `Reset the configuration file to its default state.

This will:
- Clear the current user
- Reset database URL to default
- Keep the configuration file structure intact

Use this command for testing or to start fresh.`,
	RunE: runReset,
}

func runReset(cmd *cobra.Command, args []string) error {

	defaultConfig := config.Config{
		DbURL: config.DefaultDbURL,
		CurrentUserName: "",
	}

	if err := defaultConfig.SetUser(""); err != nil {
		return fmt.Errorf("failed to reset configuration: %w", err)
	}
	
	fmt.Println("🔄 Configuration has been reset to defaults")
	fmt.Println("Run 'gator login <username>' to set up your user again")
	return nil

}