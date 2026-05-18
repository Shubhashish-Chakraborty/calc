package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:     "sub",
	Aliases: []string{"subtraction"},
	Short:   "Subtract 2 numbers",
	Long:    "Carry out subtraction operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("subtraction of %s from %s = %s.\n\n", args[1], args[0], Subtract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
