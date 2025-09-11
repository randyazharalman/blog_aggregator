package main

import (
	"fmt"
	"os"

	"github.com/randyazharalman/blog_aggregator/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}