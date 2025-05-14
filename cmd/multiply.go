package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(multiplyCmd)
}

var multiplyCmd = &cobra.Command{
	Use:   "multiply",
	Short: "Calculates the product of a list of integers",
	Long:  "The 'multiply' command multiplies a list of integer arguments together. It returns the total product of all numbers provided.",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		numbers, err := parseIntArgs(args)
		if err != nil {
			return err
		}
		for i := range len(numbers) - 1 {
			numbers[i+1] = numbers[i] * numbers[i+1]
		}
		result := numbers[len(numbers)-1]
		fmt.Printf("Result: %d", result)
		return nil
	},
}
