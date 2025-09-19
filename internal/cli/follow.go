package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/randyazharalman/blog_aggregator/internal/database"
	"github.com/spf13/cobra"
)

var followCmd = &cobra.Command{
	Use:   "follow <url>",
	Short: "Follow a feed by URL",
	Long: `Follow an existing feed by its URL.

The feed must already exist in the database (added by any user).
You cannot follow the same feed twice.`,
	Example: `  gator follow https://blog.boot.dev/index.xml
  gator follow https://www.wagslane.dev/index.xml`,
	Args: cobra.ExactArgs(1),
	RunE: middlewareLoggedIn(runFollow),
}

func runFollow(cmd *cobra.Command, args []string, user database.User) error {
	if state.Config.CurrentUserName == "" {
		fmt.Println("Error: No user logged in. Please login first with 'gator login <username>'")
		os.Exit(1)
	}

	url := strings.TrimSpace(args[0])
	if url == "" {
		return fmt.Errorf("URL cannot be empty")
	}

	user, err := state.DB.GetUser(context.Background(), state.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	feed, err := state.DB.GetFeedByURL(context.Background(), url)
		if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Error: Feed with URL '%s' not found. Use 'gator addfeed' to add it first.\n", url)
			os.Exit(1)
		}
		return fmt.Errorf("failed to get feed: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

		feedFollow, err := state.DB.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") ||
			strings.Contains(err.Error(), "duplicate key") {
			fmt.Printf("Error: You are already following the feed '%s'\n", feed.Name)
			os.Exit(1)
		}
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Printf("✅ Now following '%s' by %s\n", feedFollow.FeedName, feedFollow.UserName)
	
	return nil
}