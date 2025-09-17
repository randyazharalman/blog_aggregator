package cli

import (
	"context"
	"fmt"

	"github.com/randyazharalman/blog_aggregator/internal/rss"
	"github.com/spf13/cobra"
)

var aggCmd = &cobra.Command{
	Use:   "agg",
	Short: "Aggregate RSS feeds",
	Long: `Fetch and aggregate RSS feeds.

This command fetches RSS feeds and displays their content.
Currently fetches from https://www.wagslane.dev/index.xml for testing.`,
	RunE: runAgg,
}

func runAgg(cmd *cobra.Command, args []string) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	
	fmt.Printf("Fetching feed from: %s\n", feedURL)
	
	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}
	
	fmt.Printf("\n=== RSS Feed ===\n")
	fmt.Printf("Title: %s\n", feed.Channel.Title)
	fmt.Printf("Link: %s\n", feed.Channel.Link)
	fmt.Printf("Description: %s\n", feed.Channel.Description)
	fmt.Printf("\n=== Items (%d) ===\n", len(feed.Channel.Item))
	
	for i, item := range feed.Channel.Item {
		fmt.Printf("\n--- Item %d ---\n", i+1)
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Description: %s\n", item.Description)
		fmt.Printf("Published: %s\n", item.PubDate)
	}
	
	return nil
}