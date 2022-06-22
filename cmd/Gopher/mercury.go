package Gopher

import (
	"fmt"

	"github.com/Siddheshk02/CLI/Go-gopher/pkg/Gopher"
	"github.com/spf13/cobra"
	"github.com/kyokomi/emoji/v2"
)

var mercuryCmd = &cobra.Command{
	Use:     "Mercury",
	Aliases: []string{"mercury"},
	Short:   "Weight on Mercury",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wt := 0.00
		emoji.Println("Enter the Weight on Earth \U0001f30e : ")
		fmt.Scanln(&wt)

		res := Gopher.Mercury(wt)
		fmt.Printf("The Weight on Mercury is %f kg :)", res)
	},
}

func init() {
	rootCmd.AddCommand(mercuryCmd)
}
