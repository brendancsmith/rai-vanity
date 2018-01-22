package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "rai-vanity",
	Short:            "A vanity address generator for RaiBlocks",
	Long:             "Generate wallet seeds with desirable public addresses. Try `rai-vanity substring \"xrb\"`",
	TraverseChildren: true,
	Version:          "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
}
