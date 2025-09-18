package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/randyazharalman/blog_aggregator/internal/database"
	"github.com/randyazharalman/blog_aggregator/internal/rss"
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
	if state.Config.CurrentUserName == "" {
		fmt.Println("Error: No user logged in. Please login first with 'gator login <username>'")
		os.Exit(1)
	}
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

	fmt.Printf("Validating feed at %s...\n", url)
	_, err = rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to fetch RSS feed: %w", err)
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

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := state.DB.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		// Feed was created but follow failed - that's ok, continue
		fmt.Printf("⚠️  Feed '%s' created but failed to automatically follow it: %v\n", name, err)
	} else {
		fmt.Printf("✅ Feed '%s' created and you are now following it!\n", feedFollow.FeedName)
	}

	fmt.Printf("Feed details:\n")
	fmt.Printf("  Name: %s\n", feed.Name)
	fmt.Printf("  URL: %s\n", feed.Url)
	fmt.Printf("  Created by: %s\n", state.Config.CurrentUserName)

	return nil
}