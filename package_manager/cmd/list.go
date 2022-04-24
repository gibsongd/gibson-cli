package cmd

import (
	packagemanager "gibson/package_manager"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all assets installed in the current project.",
	Long: `List all assets installed in the current project.
	Gibson will be able to list and track only those assets installed through gibson-cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		packagemanager.ListAssets()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
