package cmd

import (
	worldparse "egrant/world-parse/internal"
	"fmt"

	"github.com/spf13/cobra"
)

// lorebookCmd represents the lorebook command
var lorebookCmd = &cobra.Command{
	Use:   "lorebook world-root-directory output-file",
	Args:  cobra.ExactArgs(2),
	Short: "Walks a worldanvil export's world directory and generates a SillyTavern lorebook from its contents",
	Long: fmt.Sprintf("%s\n%s\n%s %s\n",
		"Walks a worldanvil export's world directory and generates a SillyTavern lorebook from its contents",
		"world-rool-directory is the root of the world you want to turn into a lorebook. i.e. exportFolder/worlds/worldName",
		"Reads articles from ", worldparse.TargetDirs),
	RunE: func(cmd *cobra.Command, args []string) error {
		return worldparse.ParseWorldDirIntoLorebook(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(lorebookCmd)
}
