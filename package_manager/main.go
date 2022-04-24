package packagemanager

import (
	swagger "gibson/package_manager/client"

	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/antihax/optional"
)

var (
	httpclient *swagger.APIClient
)

func Init() {
	httpclient = swagger.NewAPIClient(swagger.NewConfiguration())
	initConfig()
	initSpinners()
}

func removeAssetFolder(assetFullName string, folder string) error {
	var assetPath string = filepath.Join(folder, assetFullName)
	return os.RemoveAll(assetPath)
}

func ClearCached(asset string) {
	startSpinner("clearing cached "+formatAsset(asset), cacheSpinner)
	if archive, cached, err := lookForCachedAsset(asset); err == nil && archive != "" {
		var assetFullName string = filepath.Join(cached.Author, cached.Title)
		if err := removeAssetFolder(assetFullName, CACHE_FOLDER); err == nil {
			removeCachedAsset(cached.Author + "/" + cached.Title)
			stopSpinner(formatAsset(cached.Author+"/"+cached.Title)+" cache cleared", cacheSpinner)
		} else {
			failSpinner("could not find cached "+formatAsset(asset), err, cacheSpinner)
		}
	} else {
		failSpinner("could not clear cached "+formatAsset(asset), err, cacheSpinner)
	}
}

func clearLocal(asset string) {
	startSpinner("removing "+formatAsset(asset), cacheSpinner)
	if err := removeAssetFolder(asset, ASSETS_FOLDER); err == nil {
		removeAsset(asset)
		stopSpinner(formatAsset(asset)+" removed", cacheSpinner)
	} else {
		failSpinner("could not remove "+formatAsset(asset), err, cacheSpinner)
	}
}

// Look for cached Asset by <author>/<name>
// 1) if there's a folder with such name, return the archive's path
// 2) if there's no archive but the Asset was cached in config, use those information to install
// 3) else return nill and fetch+install

func getArchivedAsset(assetPath string) (string, error) {
	var err error
	if _, err = os.Stat(assetPath); err == nil {
		fs, err := ioutil.ReadDir(assetPath)
		if err != nil {
			return "", err
		}
		sort.SliceStable(fs, func(i, j int) bool { return fs[i].Name() > fs[j].Name() })
		return filepath.Join(assetPath, fs[0].Name()), nil
	}
	return "", err
}

func lookForCachedAsset(asset string) (string, swagger.AssetDetails, error) {
	var assetPath string
	var cachedAsset swagger.AssetDetails
	if strings.Contains(asset, "/") {
		assetPath = filepath.Join(CACHE_FOLDER, asset)
		cachedAsset = cachedConfig.CachedAssets[asset]
	} else {
		for fullName, assetDetails := range cachedConfig.CachedAssets {
			if assetDetails.AssetId == asset {
				assetPath = filepath.Join(CACHE_FOLDER, fullName)
				cachedAsset = assetDetails
				break
			}
		}
	}
	archivedPath, err := getArchivedAsset(assetPath)
	return archivedPath, cachedAsset, err
}

func installAsset(asset swagger.AssetDetails, archive string) error {
	startSpinner("installing "+formatAsset(asset.Author+"/"+asset.Title), installSpinner)
	currentFolder, err := unzip(archive, ASSETS_FOLDER)
	if err != nil {
		return err
	}
	addAsset(currentFolder, asset)
	addCachedAsset(asset)
	stopSpinner(formatAsset(asset.Author+"/"+asset.Title)+" installed successfully!", installSpinner)
	fmt.Println()
	return nil
}

func handleCached(asset string) bool {
	startSpinner("looking for "+formatAsset(asset)+" in cache", cacheSpinner)
	if archive, cached, err := lookForCachedAsset(asset); err == nil && archive != "" {
		stopSpinner(formatAsset(asset)+" found in cache", cacheSpinner)
		installAsset(cached, archive)
		return true
	}
	stopSpinner("", cacheSpinner)
	return false
}

func InstallByAuthor(asset string, forceInstall bool) {
	if !forceInstall {
		if handleCached(asset) {
			return
		}
	}

	splAsset := strings.Split(asset, "/")
	var author string = splAsset[0]
	var assetName string = splAsset[1]

	var toDownloadAsset swagger.AssetSummary = swagger.AssetSummary{}

	// Look for asset by the author
	startSpinner("looking for "+formatAsset(asset), findSpinner)

	var query swagger.AssetsApiAssetGetOpts = swagger.AssetsApiAssetGetOpts{
		User:         optional.NewString(author),
		GodotVersion: optional.NewString("3.4"),
	}

	paginatedAssetList, response, err := httpclient.AssetsApi.AssetGet(nil, &query)
	if err != nil && response.StatusCode < 400 {
		failSpinner("fetching failed", err, findSpinner)
		return
	}

	// If there's a list, look for the asset
	if len(paginatedAssetList.Result) < 1 {
		failSpinner("couldn't find assets related to @"+author, errors.New("author doesn't exist"), findSpinner)
	}
	for _, foundAsset := range paginatedAssetList.Result {
		if foundAsset.Title == assetName {
			toDownloadAsset = foundAsset
			stopSpinner(formatAsset(asset)+" found!", findSpinner)
			break
		}
	}

	// If there's an asset, download it and install it
	if toDownloadAsset.AssetId == "" {
		failSpinner("couldn't find "+formatAsset(asset), errors.New("asset doesn't exist"), findSpinner)
		return
	}

	InstallById(toDownloadAsset.AssetId, forceInstall)
}

func InstallById(assetId string, forceInstall bool) {
	if !forceInstall {
		if handleCached(assetId) {
			return
		}
	}

	var assetDetails swagger.AssetDetails = swagger.AssetDetails{}

	startSpinner("retrieving "+formatAsset(assetId)+" info", getSpinner)

	assetDetails, response, err := httpclient.AssetsApi.AssetIdGet(nil, assetId)
	if err != nil {
		failSpinner("download failed", err, getSpinner)
		return
	}
	if response.StatusCode == 404 {
		failSpinner("download failed", errors.New("could not find asset"), getSpinner)
		return
	}

	var asset string = assetDetails.Author + "/" + assetDetails.Title
	stopSpinner(formatAsset(asset)+" info retrieved!", getSpinner)

	var assetFolder string = filepath.Join(CACHE_FOLDER, asset)
	os.MkdirAll(assetFolder, os.ModePerm)

	startSpinner("downloading "+formatAsset(asset), downloadSpinner)

	var archive string = filepath.Join(assetFolder, assetDetails.Version+".zip")
	if err := downloadFile(archive, assetDetails.DownloadUrl); err != nil {
		failSpinner("download failed", err, downloadSpinner)
		return
	}
	stopSpinner(formatAsset(asset)+" downloaded!", downloadSpinner)

	if err := installAsset(assetDetails, archive); err != nil {
		failSpinner("could not install "+formatAsset(asset), err, installSpinner)
	}
}

func InstallByConfig(forceInstall bool) {
	for _, gibsonAsset := range projectConfig.Assets {
		InstallById(gibsonAsset.Id, forceInstall)
	}
}

func InstallAsset(asset string, forceInstall bool) {
	if strings.Contains(asset, "/") {
		InstallByAuthor(asset, forceInstall)
	} else {
		InstallById(asset, forceInstall)
	}
}

func uninstallAsset(asset string) bool {
	for currentFolder, gibsonAsset := range projectConfig.Assets {
		if gibsonAsset.Id == asset || gibsonAsset.FullName == asset {
			clearLocal(currentFolder)
			return true
		}
	}
	return false
}

func UninstallAsset(asset string, clearClached bool) {
	startSpinner("uninstalling "+formatAsset(asset), findSpinner)
	if clearClached {
		ClearCached(asset)
	}
	if uninstallAsset(asset) {
		stopSpinner(formatAsset(asset)+" uninstalled", findSpinner)
	} else {
		failSpinner("couldn't find "+formatAsset(asset), errors.New("Maybe it was already deleted or misstyped."), findSpinner)
	}
}

func ListAssets() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	var i int = 0
	fmt.Fprintln(writer, "assets/")
	for folder, asset := range projectConfig.Assets {
		var prefix string = "├─ "
		if i == len(projectConfig.Assets)-1 {
			prefix = "└─ "
		}
		var line []string = []string{prefix + string(folder) + "/", asset.Id, "(" + asset.FullName + ")"}
		fmt.Fprintln(writer, strings.Join(line, "\t"))
		i++
	}
	fmt.Fprintln(writer)
	writer.Flush()
}

func SearchAsset(search string, type_ string, category string, support string, user string, godotVersion string,
	maxResults int16, page int16, offset int16, sort string, reverse bool) {
	var query swagger.AssetsApiAssetGetOpts = swagger.AssetsApiAssetGetOpts{
		Filter:       optional.NewString(search),
		GodotVersion: optional.NewString(godotVersion),
		Sort:         optional.NewString(sort),
	}
	if type_ != "" {
		query.Type_ = optional.NewString(type_)
	}
	if category != "" {
		query.Category = optional.NewString(category)
	}
	if support != "" {
		query.Support = optional.NewString(support)
	}
	if user != "" {
		query.User = optional.NewString(user)
	}
	if maxResults != 0 {
		query.MaxResults = optional.NewString(fmt.Sprint(maxResults))
	}
	if page != 0 {
		query.Page = optional.NewString(fmt.Sprint(page))
	}
	if offset != 0 {
		query.Offset = optional.NewString(fmt.Sprint(offset))
	}
	if reverse {
		query.Reverse = optional.NewBool(reverse)
	}

	if paginatedAssetList, _, err := httpclient.AssetsApi.AssetGet(nil, &query); err != nil {
		return
	} else {
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
		fmt.Fprintln(writer, "[id]\t[user]\t[title]")
		for _, assetSummary := range paginatedAssetList.Result {
			fmt.Fprintln(writer, strings.Join([]string{assetSummary.AssetId, assetSummary.Author, assetSummary.Title}, "\t"))
		}
		fmt.Fprintln(writer)
		writer.Flush()
	}
}
