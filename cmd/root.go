package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calc",
	Short: "calc is a cli tool for performing basic mathematical operations",
	Long:  "calc is a cli tool for performing basic mathematical operations - addition, multiplication, division and subtraction.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calc Invoked!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing calc '%s'\n", err)
		os.Exit(1)
	}
}
