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

var addFeedCmd = &cobra.Command{
	Use:   "addfeed <name> <url>",
	Short: "Add a new feed",
	Long: `Add a new feed in the database.

This will create a new feed with the specified name and url.
The url must be unique.`,
	Args: cobra.ExactArgs(2),
	RunE: runAddFeed,
}

func runAddFeed(cmd *cobra.Command, args []string) error {
	name:= strings.TrimSpace(args[0])
	url:= strings.TrimSpace(args[1])

	if name == "" {
		return fmt.Errorf("username cannot be empty")
	}
	if url == "" {
		return fmt.Errorf("url cannot be empty")
	}

	currentUser, err := state.DB.GetUser(context.Background(), state.Config.CurrentUserName)
	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		Name: name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url: url,
		UserID: currentUser.ID ,
	}

	feed, err := state.DB.CreateFeed(context.Background(), feedParams) 
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "duplicate key") {
			fmt.Printf("Error: URL '%s' already exists\n", url)
			os.Exit(1)
		}
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Printf("✅ Feed '%s' has been created!\n", feed.Name)
	fmt.Printf("Feed details:\n")
	fmt.Printf("  ID: %s\n", feed.ID)
	fmt.Printf("  Name: %s\n", feed.Name)
	fmt.Printf("  URL: %s\n", feed.Url)
	fmt.Printf("  User ID: %s\n", feed.UserID)
	fmt.Printf("  Created: %s\n", feed.CreatedAt.Format(time.RFC3339))
	return nil
}