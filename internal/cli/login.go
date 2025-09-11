package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login <username>",
	Short: "Login with a username",
	Long: `Set the current user in the configuration file.

This will update your ~/.gatorconfig.json file with the specified username.
The username will be used to track your personal feed subscriptions.`,
	Example: `  gator login alice
  gator login john_doe`,
	Args: cobra.ExactArgs(1),
	RunE: runLogin,
}

func runLogin(cmd *cobra.Command, args []string) error {
	username := strings.TrimSpace(args[0])

	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	if err := state.Config.SetUser(username); err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	fmt.Printf("✅ User has been set to: %s\n", username)
	return nil
}