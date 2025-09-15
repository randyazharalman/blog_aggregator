package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/randyazharalman/blog_aggregator/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}