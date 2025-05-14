package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(divideCmd)
}

var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "Divides the first number by the remaining, in order",
	Long:  "The 'divide' command divides the first number by the following numbers in sequence. It performs integer division and returns the result. If any divisor is 0, the operation is aborted to prevent division by zero.",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		numbers, err := parseIntArgs(args)
		if err != nil {
			return err
		}

		for i := range len(numbers) - 1 {
			if numbers[i+1] == 0 {
				return fmt.Errorf("cannot divide by zero")
			}
			numbers[i+1] = numbers[i] / numbers[i+1]
		}
		result := numbers[len(numbers)-1]
		fmt.Printf("Result: %d", result)
		return nil
	},
}
