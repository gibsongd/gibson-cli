package cmd

import (
	packagemanager "gibson/package_manager"

	"github.com/spf13/cobra"
)

var (
	clearFlag bool
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Aliases: []string{"u"},
	Short:   "Uninstall an asset from the current project",
	Long: `Uninstall an asset from the current project.
	The asset can be passed both by ID or by FullNmae ({author}/{name}).`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		packagemanager.UninstallAsset(args[0], clearFlag)
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
	uninstallCmd.Flags().BoolVarP(&clearFlag, "clear", "c", false, "Clear all the cached versions of the asset you want to uninstall.")
}
