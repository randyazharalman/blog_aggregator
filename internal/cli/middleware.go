package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/randyazharalman/blog_aggregator/internal/database"
	"github.com/spf13/cobra"
)

func middlewareLoggedIn(handler func(*cobra.Command, []string, database.User) error )func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if state.Config.CurrentUserName == "" {
			fmt.Println("Error: No user logged in. Please login first with 'gator login <username>'")
			os.Exit(1)
		}

		user, err := state.DB.GetUser(context.Background(), state.Config.CurrentUserName)
		if err != nil {
			fmt.Printf("Error: Failed to get current user: %v\n", err)
			os.Exit(1)
		}
		
	return handler(cmd, args, user)
  }
}