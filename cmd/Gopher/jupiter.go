package Gopher

import (
	"fmt"

	"github.com/Siddheshk02/CLI/Go-gopher/pkg/Gopher"
	"github.com/spf13/cobra"
	"github.com/kyokomi/emoji/v2"
)

var jupiterCmd = &cobra.Command{
	Use:     "Jupiter",
	Aliases: []string{"jupiter"},
	Short:   "Weight on Jupiter",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wt := 0.00
		emoji.Println("Enter the Weight on Earth \U0001f30e : ")
		fmt.Scanln(&wt)

		res := Gopher.Jupiter(wt)
		fmt.Printf("The Weight on Jupiter is %f kg :)", res)
	},
}

func init() {
	rootCmd.AddCommand(jupiterCmd)
}
