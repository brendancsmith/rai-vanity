package cmd

import (
	"fmt"
	"os"
	"strings"

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

		substring := args[0]
		index = cmd.Flags().GetString("index")
		count = cmd.Flags().GetInt("count")

		run(args[0], atoi(index), cmd.Flags)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// TODO: check if index flag is valid integer (within range) or "*"

		return nil
	},
}

func init() {
	rootCmd.AddCommand(substringCmd)

	substringCmd.Flags().StringP("index", "i", "5", "The index in the address to match the substring.\n"+
		"\t\"*\" will match the substring at any position in the address.\n"+
		"\t\"-1\" will match at the end of the address.")

	substringCmd.Flags().IntP("count", "c", 1, "Number of valid addresses to generate before exiting, or 0 for infinite.")

}

func run(string substring, int index, int count) {
	iterations := estimatedIterations(substring)

	fmt.Println("Estimated number of iterations needed:", iterations)
	for i := 0; i < args.count || args.count == 0; i++ {
		seed, addr, err := generateVanityAddress(substring)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Found matching address!")
		fmt.Printf("Seed: %s\n", strings.ToUpper(seed))
		fmt.Printf("Address: %s\n", addr)
	}
}
