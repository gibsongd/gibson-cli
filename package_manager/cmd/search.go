package cmd

import (
	"errors"
	"fmt"
	packagemanager "gibson/package_manager"
	"strings"

	"github.com/spf13/cobra"
)

var validTypes []string = []string{"any", "project", "addon"}
var validSupports []string = []string{"", "official", "community", "testing"}
var validSorts []string = []string{"", "rating", "cost", "name", "updated"}

var (
	typeFlag         string
	categoryFlag     string
	supportFlag      string
	filterFlag       string
	userFlag         string
	godotVersionFlag string
	maxResultsFlag   int16
	pageFlag         int16
	offsetFlag       int16
	sortFlag         string
	reverseFlag      bool
)

func invalidFlag(flagName string, validValues []string) error {
	return errors.New(fmt.Sprintf("Invalid value for '%s' flag, must be one of %s", flagName, validValues))
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Look for an asset in the asset library.",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Must search at least a word!")
		} else {
			filterFlag = strings.Join(args, " ")
		}

		if !packagemanager.Contains(validTypes, typeFlag) {
			return invalidFlag("type", validTypes)
		}

		if !packagemanager.Contains(validSupports, supportFlag) {
			return invalidFlag("support", validSupports)
		}

		if !packagemanager.Contains(validSorts, sortFlag) {
			return invalidFlag("sort", validSorts)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		packagemanager.SearchAsset(
			filterFlag, typeFlag, categoryFlag, supportFlag, userFlag, godotVersionFlag,
			maxResultsFlag, pageFlag, offsetFlag, sortFlag, reverseFlag,
		)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&typeFlag, "type", "t", "any", fmt.Sprintf(`The asset's type, can be one of %s.`, validTypes))
	searchCmd.Flags().StringVarP(&categoryFlag, "category", "c", "", `The category the asset belongs to.`)
	searchCmd.Flags().StringVarP(&supportFlag, "support", "s", "", fmt.Sprintf(`The asset's support level, can be one of %s.`, validSupports))
	searchCmd.Flags().StringVarP(&userFlag, "user", "u", "", `The author's username.`)
	searchCmd.Flags().StringVarP(&godotVersionFlag, "godotVersion", "g", "3.4", "The Godot version the asset's latest version is intended for (in 'major.minor' format).\nThis field is present for compatibility reasons with the Godot editor. See also the 'versions' array.")
	searchCmd.Flags().Int16VarP(&maxResultsFlag, "maxResults", "m", 0, `The maximum number of results to display.`)
	searchCmd.Flags().Int16VarP(&pageFlag, "page", "p", 0, `Pages to skip`)
	searchCmd.Flags().Int16VarP(&offsetFlag, "offset", "o", 0, `Rows to skip`)
	searchCmd.Flags().StringVarP(&sortFlag, "sort", "S", "", fmt.Sprintf(`The order applied to the results, can be one of %s.`, validSorts))
	searchCmd.Flags().BoolVarP(&reverseFlag, "revers", "r", false, `If the result list should be reversed or not.`)

}
