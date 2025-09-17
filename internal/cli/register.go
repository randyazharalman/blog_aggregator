package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/randyazharalman/blog_aggregator/internal/database"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register <username>",
	Short: "Register a new user",
	Long: `Register a new user in the database.

This will create a new user account with the specified username.
The username must be unique - you cannot register with an existing username.

After successful registration, you will be automatically logged in.`,
	Example: `  gator register alice
  gator register john_doe`,
	Args: cobra.ExactArgs(1),
	RunE: runRegister,
}

func runRegister(cmd *cobra.Command, args []string) error {
	name:= strings.TrimSpace(args[0])

	if name == "" {
		return fmt.Errorf("username cannot be empty")
	}

	userParams := database.CreateUserParams{
		ID: uuid.New(),
		Name: name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := state.DB.CreateUser(context.Background(), userParams) 
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "duplicate key") {
			fmt.Printf("Error: User '%s' already exists\n", name)
			os.Exit(1)
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := state.Config.SetUser(name); err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}
	fmt.Printf("✅ User '%s' has been created and set as current user!\n", name)
	fmt.Printf("User details:\n")
	fmt.Printf("  ID: %s\n", user.ID)
	fmt.Printf("  Name: %s\n", user.Name)
	fmt.Printf("  Created: %s\n", user.CreatedAt.Format(time.RFC3339))
	return nil
}