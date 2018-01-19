package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rai-vanity",
	Short: "A vanity address generator for RaiBlocks",
	// Long: `A longer description that spans multiple lines
	// 		 and likely contains examples and usage of using your application.`

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },

	// By default Cobra only parses local flags on the target command.
	// By enabling this, Cobra will parse local flags on each command
	// before executing the target command.
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("regex cmd called")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
}
