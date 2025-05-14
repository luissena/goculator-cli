package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "goculator-cli",
		Short: "A simple calculator CLI",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
}

func parseIntArgs(args []string) ([]int, error) {
	var numbers []int

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("invalid argument %s, expected integer", arg)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
