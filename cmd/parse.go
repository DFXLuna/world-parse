package cmd

import (
	worldparse "egrant/world-parse/internal"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ErrBadPath = errors.New("error with path")
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse file",
	Short: "parses a single file and outputs the result to stdout",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		if worldparse.Exists(args[0]) {
			content, err := worldparse.ParseOneFile(args[0])
			fmt.Printf("%v", content)
			return err
		}
		return ErrBadPath
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
