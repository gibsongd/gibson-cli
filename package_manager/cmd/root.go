package cmd

import (
	packagemanager "gibson/package_manager"
	"os"

	"github.com/spf13/cobra"
)

var PackageManager = rootCmd

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "package_manager",
	Aliases: []string{"pm"},
	Short:   "The Gibson's Package Manager.",
	Long:    `Gibson's Package Manager can be used to manage Godot Assets directly from cli.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	packagemanager.Init()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
