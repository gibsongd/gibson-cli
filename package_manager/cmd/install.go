package cmd

import (
	"errors"
	packagemanager "gibson/package_manager"
	"strings"

	"github.com/spf13/cobra"
)

var (
	asset           string
	assets          []string
	isFullName      bool
	forceFlag       bool
	installByConfig bool
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     `install`,
	Aliases: []string{"i"},
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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			installByConfig = true
		} else {
			if len(args) == 1 {
				asset = args[0]
				if isFullName = strings.Contains(asset, "/"); isFullName {
					if spl := strings.Split(asset, "/"); spl[0] == "" || spl[1] == "" {
						return errors.New("\033[31mInvalid asset, must match {author}/{name} format!\033[0m")
					}
				}
			} else {
				assets = args
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if installByConfig {
			packagemanager.InstallByConfig(forceFlag)
			return
		}

		if len(assets) > 0 {
			for _, asset := range assets {
				packagemanager.InstallAsset(asset, forceFlag)
			}
			return
		}
		if isFullName {
			packagemanager.InstallByAuthor(asset, forceFlag)
			return
		} else {
			packagemanager.InstallById(asset, forceFlag)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force install the asset, bypassing the cache and overwriting it.")
}
