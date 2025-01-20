package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zero",
	Short: "zero is a cli tool for math.",
	Long:  "zero is a cli tool for math - adding, multiplication, subtraction, and division.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops.")
		os.Exit(1)
	}
}
