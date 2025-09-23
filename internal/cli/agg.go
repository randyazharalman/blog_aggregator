package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var aggCmd = &cobra.Command{
	Use:   "agg <time_between_reqs>",
	Short: "Aggregate RSS feeds continuously",
	Long: `Continuously fetch and display RSS feeds in a loop.

The time_between_reqs parameter specifies how long to wait between
fetching feeds. It should be a duration string like "1m", "30s", "1h", etc.

This command will run indefinitely until stopped with Ctrl+C.
It fetches feeds in order of when they were last fetched (oldest first).

Examples of duration formats:
  - 30s (30 seconds)
  - 1m (1 minute) 
  - 5m (5 minutes)
  - 1h (1 hour)`,
	Example: `  gator agg 1m
  gator agg 30s
  gator agg 2m30s`,
	Args: cobra.ExactArgs(1),
	RunE: runAgg,
}

func runAgg(cmd *cobra.Command, args []string) error {
	// Parse the duration argument
	timeBetweenRequests, err := time.ParseDuration(args[0])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	fmt.Println("Press Ctrl+C to stop...")
	fmt.Println()

	// Set up graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nShutting down aggregator...")
		cancel()
	}()

	// Create ticker for periodic fetching
	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	// Run immediately, then on each tick
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Aggregator stopped")
			return nil
		default:
			// Scrape feeds
			if err := scrapeFeeds(); err != nil {
				fmt.Printf("Error scraping feeds: %v\n", err)
			}
		}

		// Wait for next tick or context cancellation
		select {
		case <-ctx.Done():
			fmt.Println("Aggregator stopped")
			return nil
		case <-ticker.C:
			// Continue to next iteration
		}
	}
}