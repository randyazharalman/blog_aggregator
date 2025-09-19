package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/randyazharalman/blog_aggregator/internal/database"
	"github.com/spf13/cobra"
)

var followingCmd = &cobra.Command{
	Use:   "following",
	Short: "List feeds you are following",
	Long: `Display all feeds that the current user is following.

Shows the names of the feeds in alphabetical order.
You must be logged in to use this command.`,
	RunE: middlewareLoggedIn(runFollowing),
}

func runFollowing(cmd *cobra.Command, args []string, user database.User) error {
	if state.Config.CurrentUserName == "" {
		fmt.Println("Error: No user logged in. Please login first with 'gator login <username>'")
		os.Exit(1)
	}

	user, err := state.DB.GetUser(context.Background(), state.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	feedFollows, err := state.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Printf("You are not following any feeds yet.\n")
		fmt.Printf("Use 'gator follow <url>' to follow a feed.\n")
		return nil
	}

	fmt.Printf("Feeds you are following:\n")
	for _, feedFollow := range feedFollows {
		fmt.Printf("* %s\n", feedFollow.FeedName)
	}

	return nil
}