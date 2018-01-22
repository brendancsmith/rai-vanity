package cmd

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/brendancsmith/rai-vanity/rai"
	"github.com/spf13/cobra"
)

var (
	flagIndex string
	flagCount int

	optWildcardIndex bool
	optIndex         int
	optCount         int
	argSubstring     string
)

// substringCmd represents the substring command
var substringCmd = &cobra.Command{
	Use:   "substring",
	Short: "Search for an address containing a substring",
	// Long:
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("substring cmd called")
		findSubstring(argSubstring, optIndex, optWildcardIndex, optCount)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		argSubstring = args[0]

		if flagIndex == "*" {
			optWildcardIndex = true
			optIndex = 0
		} else {
			optWildcardIndex = false
			var err error
			optIndex, err = strconv.Atoi(flagIndex)
			if err != nil {
				return err
			}

			if optIndex == -1 {
				optIndex = 64 - len(argSubstring)
			}
		}

		optCount = int(math.Max(0, float64(flagCount)))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(substringCmd)

	substringCmd.Flags().StringVarP(&flagIndex, "index", "i", "5", "The index in the address to match the substring.\n"+
		"\t\"*\" will match the substring at any position in the address.\n"+
		"\t\"-1\" will match at the end of the address.")

	substringCmd.Flags().IntVarP(&flagCount, "count", "c", 1, "Number of valid addresses to generate before exiting, or 0 for infinite.")
}

func findSubstring(substring string, index int, wildcardIndex bool, count int) {
	iterations := rai.EstimateIterations(substring, index)

	fmt.Println("Count:", count)

	fmt.Println("Estimated number of iterations needed:", iterations)
	for i := 0; i < count || count == 0; i++ {
		seed, addr, err := rai.GenerateVanityAddress(substring, index, iterations)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Found matching address!")
		fmt.Printf("Seed: %s\n", strings.ToUpper(seed))
		fmt.Printf("Address: %s\n", addr)
	}
}
