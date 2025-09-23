package cli

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/randyazharalman/blog_aggregator/internal/rss"
)

func scrapeFeeds() error {
	feed, err := state.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No feed available to fetch")
			return nil
		}
		return fmt.Errorf("failed to get next feed: %w", err)
	}
	
	
	err = state.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err !=  nil {
		return fmt.Errorf("failed to mark feed as fetched: %w", err)
	}
	
	fmt.Printf("Fetching feed: %s from %ss\n", feed.Name, feed.Url)
	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		 fmt.Printf("Error fetching feed %s: %v",feed.Name,  err)
		return nil
	}

	fmt.Printf("Found %d posts in %s:\n", len(rssFeed.Channel.Item), feed.Name)
	
	for i, item := range rssFeed.Channel.Item {
		fmt.Printf("  %d. %s\n", i+1, item.Title)
		if item.Link != "" {
			fmt.Printf("     URL: %s\n", item.Link)
		}
		if item.PubDate != "" {
			fmt.Printf("     Published: %s\n", item.PubDate)
		}
		fmt.Println()
	}
	
	fmt.Printf("Completed fetching %s\n", feed.Name)
	fmt.Println("---")

	return nil
}