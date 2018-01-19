package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// substringCmd represents the substring command
var substringCmd = &cobra.Command{
	Use:   "substring",
	Short: "Search for an address containing a substring",
	// Long:
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("substring cmd called")
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// TODO: check if index flag is valid integer (within range) or "*"

		return nil
	},
}

func init() {
	rootCmd.AddCommand(substringCmd)

	substringCmd.Flags().StringP("index", "i", "*", "The index in the address to match the substring.\n"+
		"\t\"*\" will match the substring at any position in the address.\n"+
		"\t\"-1\" will match at the end of the address.")
}
