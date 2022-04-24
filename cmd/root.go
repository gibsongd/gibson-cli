package cmd

import (
	cmd "gibson/package_manager/cmd"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gibson",
	Short: "The Gibson CLI.",
	Long:  ``,
}

func AddCommand(cmds *cobra.Command) {
	rootCmd.AddCommand(cmds)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmd.PackageManager)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
