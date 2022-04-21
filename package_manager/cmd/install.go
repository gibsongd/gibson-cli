/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	packagemanager "gibson/package_manager"

	"github.com/spf13/cobra"
)

var (
	asset   string
	assetId string
	force   bool
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     `install`,
	Aliases: []string{"i"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
			return
		}

		if assetId != "" {
			packagemanager.InstallById(assetId, force)
			return
		}

		if asset != "" {
			packagemanager.InstallByAuthor(asset, force)
			return
		}

		if packagemanager.Contains(args, ".") {
			packagemanager.InstallByConfig(force)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	installCmd.Flags().StringVar(&assetId, "id", "", "Install an asset by its id.")
	installCmd.Flags().StringVar(&asset, "asset", "", "Install an asset by its {author}/{name} combination.")
	installCmd.Flags().BoolVar(&force, "f", false, "Force install the asset, bypassing the cache and overwriting it.")
}
