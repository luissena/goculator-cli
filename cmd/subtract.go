package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(subtractCmd)
}

var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "Calculates the subtraction of a list of integers",
	Long:  "The 'subtract' command subtracts a list of integer arguments in order. It starts with the first number and subtracts the rest sequentially.",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		numbers, err := parseIntArgs(args)
		if err != nil {
			return err
		}

		var result int

		for _, v := range numbers {
			result -= v
		}

		fmt.Printf("Result: %d", result)

		return nil
	},
}
