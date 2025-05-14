package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sumCmd)
}

var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "Calculates the sum of a list of integers",
	Long:  "The 'sum' command receives a list of integer arguments and calculates their total. It validates each input, performs the addition, and prints the result to the standard output.",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		numbers, err := parseIntArgs(args)
		if err != nil {
			return err
		}

		var result int

		for _, v := range numbers {
			result += v
		}

		fmt.Printf("Result: %d", result)

		return nil
	},
}
