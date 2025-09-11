package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Version = "1.0.0"
	AppName = "gator"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Display the current version of the gator CLI tool.`,
	Run:   runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("%s v%s\n", AppName, Version)
	fmt.Println("A blog aggregator CLI tool")
}