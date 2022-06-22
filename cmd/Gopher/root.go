package Gopher

import (
	"fmt"
	"os"

	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{

	Use:   "Gopher",
	Short: "Gopher - A simple CLI to Check weight on different planets",
	Long: `Gopher - A simple CLI
- One can use to Check weight on different planets`,

	Run: func(cmd *cobra.Command, args []string) {
		emoji.Println("\U0001f680 \U0001f680 FIND OUT THE WEIGHT ON DIFFERENT PLANETS \U0001f680 \U0001f680")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI :( '%s'\n\n", err)
		os.Exit(1)
	}
}
