package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var feedsCmd = &cobra.Command{
	Use:   "feeds",
	Short: "List all feeds",
	Long: `Display a list of all feeds in the database.`,
	RunE: runFeeds,
}

func runFeeds(cmd *cobra.Command, args []string) error {
	feeds, err := state.DB.GetFeedsWithUser(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}
	
	if len(feeds) == 0 {
		fmt.Println("No feeds found. Use 'addfeed <name> <url>' to create a feed.")
		return nil
	}
	
	
	for i, feed := range feeds {
			fmt.Printf("%d. Feed Name: %s\n", i + 1, feed.Name)
			fmt.Printf("   Feed URL: %s\n", feed.Url)
			fmt.Printf("   User: %s\n", feed.Name_2)
	}
	
	return nil
}