package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/randyazharalman/blog_aggregator/internal/database"
	"github.com/spf13/cobra"
)

var unfollowCmd = &cobra.Command{
	Use:   "unfollow <url>",
	Short: "Unfollow a feed by URL",
	Long: `Unfollow a feed by its URL.

This will remove your subscription to the feed. The feed will remain
in the database and other users can still follow it.

You must be currently following the feed to unfollow it.`,
	Example: `  gator unfollow https://blog.boot.dev/index.xml
  gator unfollow https://www.wagslane.dev/index.xml`,
	Args: cobra.ExactArgs(1),
	RunE: middlewareLoggedIn(runUnfollow),
}

func runUnfollow(cmd *cobra.Command, args []string, user database.User) error {
	url := strings.TrimSpace(args[0])
	if url == "" {
		return fmt.Errorf("URL cannot be empty")
	}

	feed, err := state.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Error: Feed with URL '%s' not found\n", url)
			os.Exit(1)
		}
		return fmt.Errorf("failed to get feed: %w", err)
	}

	params := database.GetFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, err = state.DB.GetFeedFollow(context.Background(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Error: You are not following the feed '%s'\n", feed.Name)
			os.Exit(1)
		}
		return fmt.Errorf("failed to check feed follow status: %w", err)
	}

	deleteParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = state.DB.DeleteFeedFollow(context.Background(), deleteParams)
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %w", err)
	}

	fmt.Printf("✅ Successfully unfollowed '%s'\n", feed.Name)
	return nil
}