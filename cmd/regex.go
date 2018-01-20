package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// regexCmd represents the regex command
var regexCmd = &cobra.Command{
	Use:   "regex",
	Short: "Find addresses that match a regular expression (regex)",
	// Long
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("regex cmd called")
	},
}

func init() {
	rootCmd.AddCommand(regexCmd)

	substringCmd.Flags().IntP("count", "c", 1, "Number of valid addresses to generate before exiting, or 0 for infinite.")
}
