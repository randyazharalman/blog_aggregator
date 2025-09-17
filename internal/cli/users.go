package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "List all registered users",
	Long: `Display a list of all users registered in the database.

The currently logged-in user (if any) will be marked with "(current)".
Users are displayed in alphabetical order.`,
	RunE: runUsers,
}

func runUsers(cmd *cobra.Command, args []string) error {
	users, err := state.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	
	if len(users) == 0 {
		fmt.Println("No users found. Use 'gator register <username>' to create a user.")
		return nil
	}
	
	currentUser := state.Config.CurrentUserName
	
	for _, user := range users {
		if user.Name == currentUser && currentUser != "" {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	
	return nil
}