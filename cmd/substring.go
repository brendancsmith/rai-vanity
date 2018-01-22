package cmd

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/brendancsmith/rai-vanity/app"
	"github.com/frankh/rai/address"
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

		var predicate app.StringPredicate

		if optWildcardIndex {
			predicate = func(address string) bool {
				return strings.Contains(address, argSubstring)
			}
		} else {
			length := len(argSubstring)
			predicate = func(address string) bool {
				return (address[optIndex:optIndex+length] == argSubstring)
			}
		}

		// TODO: cancel on count
		ctx := context.Background()

		ch, err := app.GenerateSeeds(ctx, predicate)
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			seed := <-ch

			pub, _ := address.KeypairFromSeed(seed, 0)
			account := address.PubKeyToAddress(pub)

			if !address.ValidateAddress(account) {
				// TODO: panic here?
				log.Fatal("Address generated had an invalid checksum!\nPlease create an issue on github: https://github.com/frankh/rai-vanity")
			}

			log.Println("Found matching address!")
			log.Printf("Seed: %s\n", strings.ToUpper(seed))
			log.Printf("Address: %s\n", account)
		}()
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
