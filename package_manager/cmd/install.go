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
	Args:    cobra.MinimumNArgs(1),
	Short:   "Install an asset from the AssetLibrary",
	Long: `Install an asset from the AssetLibrary into the current project.
	
	Everytime an asset is installed via gibson, the installed addon will be cached.
	When installing an asset, gibson will look for the asset author, name and id in the local cache,
	and if a matching asset is found, it will be installed in the current project.

	Assets can also be installed in batch using a gibson.json file.
	When installing an addon using gibson, a gibson.json file will be created at the root of the current project,
	and it will be used to store the information related to the addon installed.
	You can share the gibson.json file instead of the whole asset/ folder of your projet to minimize the size of your project,
	when sharing it with your team or publically.
	Assets registered in the gibson.json file can be installed via the < gibson pm install . > command.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if assetId != "" {
			packagemanager.InstallById(assetId, force)
			return
		}

		if asset != "" {
			packagemanager.InstallByAuthor(assetId, force)
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
	installCmd.Flags().StringVar(&assetId, "id", "", "Install an asset by its id.")
	installCmd.Flags().StringVar(&asset, "asset", "", "Install an asset by its {author}/{name} combination.")
	installCmd.Flags().BoolVarP(&force, "force", "f", false, "Force install the asset, bypassing the cache and overwriting it.")
}
