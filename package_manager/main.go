package packagemanager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"
)

func Init() {
	initConfig()
	initSpinners()
}

func clearAsset(asset string, folder string) error {
	var assetPath string = filepath.Join(folder, asset)
	fmt.Println(assetPath)
	return os.RemoveAll(assetPath)
}

func clearCache(asset string) {
	startSpinner("clearing cached "+formatAsset(asset), cacheSpinner)
	if archive, cached, err := lookForCachedAsset(asset); err == nil && archive != "" {
		if err := clearAsset(filepath.Join(cached.Author, cached.Title), cacheFolder); err == nil {
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
	if err := clearAsset(asset, addonsFolder); err == nil {
		removeAddon(Folder(asset))
		stopSpinner(formatAsset(asset)+" removed", cacheSpinner)
	} else {
		failSpinner("could not remove "+formatAsset(asset), err, cacheSpinner)
	}
}

// Look for cached addon by <author>/<name>
// 1) if there's a folder with such name, return the archive's path
// 2) if there's no archive but the addon was cached in config, use those information to install
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

func lookForCachedAsset(asset string) (string, Addon, error) {
	var assetPath string
	var cachedAddon Addon
	if strings.Contains(asset, "/") {
		assetPath = filepath.Join(cacheFolder, asset)
		cachedAddon = cachedConfig.CachedAddons[AssetName(asset)]
	} else {
		for assetName, addon := range cachedConfig.CachedAddons {
			if addon.AssetId == asset {
				assetPath = filepath.Join(cacheFolder, string(assetName))
				cachedAddon = addon
				break
			}
		}
	}
	archivedPath, err := getArchivedAsset(assetPath)
	return archivedPath, cachedAddon, err
}

func installAsset(addon Addon, archive string) {
	var asset string = addon.Author + "/" + addon.Title
	startSpinner("installing "+formatAsset(asset), installSpinner)
	folder, err := unzip(archive, addonsFolder)
	if err != nil {
		failSpinner("could not install "+formatAsset(asset), err, installSpinner)
		return
	}

	addAddon(Folder(folder), GibsonAddon{Asset: asset, Id: addon.AssetId})
	addCachedAsset(asset, addon)

	stopSpinner(formatAsset(asset)+" installed successfully!", installSpinner)
}

func installByAuthor(asset string, clearCached bool) {
	tempPack := strings.Split(asset, "/")
	var author string = tempPack[0]
	var assetName string = tempPack[1]

	var toDownloadAddon Addon = Addon{}
	var assetResult AssetResult = AssetResult{}

	// Look for asset by the author
	startSpinner("looking for "+formatAsset(asset), findSpinner)

	err := doGet("/asset?user="+author+"&godot_version=3.4", &assetResult)
	if err != nil {
		failSpinner("fetching failed", err, findSpinner)
		return
	}

	// If there's a list, look for the asset
	if len(assetResult.Result) < 1 {
		failSpinner("couldn't find assets related to @"+author, errors.New("author doesn't exist"), findSpinner)
	}
	for _, addon := range assetResult.Result {
		if addon.Title == assetName {
			toDownloadAddon = addon
			spinnerMessage(formatAsset(asset)+" found!", findSpinner)
			break
		}
	}

	// If there's an asset, download it and install it
	if toDownloadAddon == (Addon{}) {
		failSpinner("couldn't find "+formatAsset(asset), errors.New("addon doesn't exist"), findSpinner)
		return
	}
	stopSpinner(formatAsset(asset)+" fetched!", findSpinner)

	installById(toDownloadAddon.AssetId)

}

func installById(assetId string) {
	var toDownloadAddon Addon = Addon{AssetId: assetId}

	startSpinner("retrieving "+formatAsset(assetId)+" info", getSpinner)

	if err := doGet("/asset/"+assetId, &toDownloadAddon); err != nil {
		failSpinner("could not find "+formatAsset(assetId), err, getSpinner)
		return
	}
	stopSpinner(formatAsset(assetId)+" info retrieved!", getSpinner)

	var asset string = toDownloadAddon.Author + "/" + toDownloadAddon.Title
	var assetFolder string = filepath.Join(cacheFolder, asset)
	os.MkdirAll(assetFolder, os.ModePerm)

	startSpinner("downloading "+formatAsset(asset), downloadSpinner)

	var archive string = filepath.Join(assetFolder, toDownloadAddon.Version+"_"+toDownloadAddon.DownloadCommit+".zip")

	if err := downloadFile(archive, toDownloadAddon.DownloadUrl); err != nil {
		failSpinner("download failed", err, downloadSpinner)
		return
	}
	stopSpinner(formatAsset(asset)+" downloaded!", downloadSpinner)

	installAsset(toDownloadAddon, archive)
}

func uninstallAsset(asset string) {
	for folder, addon := range projectConfig.Addons {
		if addon.Asset == asset || addon.Id == asset {
			clearLocal(string(folder))
			stopSpinner(formatAsset(asset)+" uninstalled", findSpinner)
			return
		}
	}
	failSpinner("couldn't find "+formatAsset(asset), errors.New("Maybe it was already deleted or misstyped."), findSpinner)
}

func InstallByConfig(clearCached bool) {
	for _, addon := range projectConfig.Addons {
		installByAuthor(addon.Asset, clearCached)
	}
}

func InstallAddon(asset string, clearCached bool) {
	if clearCached {
		clearCache(asset)
	} else {
		startSpinner("looking for "+formatAsset(asset)+" in cache", cacheSpinner)
		archive, cached, err := lookForCachedAsset(asset)
		if err == nil && archive != "" {
			stopSpinner(formatAsset(asset)+" found in cache", cacheSpinner)
			installAsset(cached, archive)
			return
		} else {
			stopSpinner("", cacheSpinner)
		}
	}

	if strings.Contains(asset, "/") {
		installByAuthor(asset, clearCached)
	} else {
		installById(asset)
	}
}

func UninstallAddon(asset string, clearCached bool) {
	startSpinner("uninstalling "+formatAsset(asset), findSpinner)
	if clearCached {
		clearCache(asset)
	}
	uninstallAsset(asset)
}

func ListAddons() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	var i int = 0
	fmt.Fprintln(writer, "addons/")
	for folder, addon := range projectConfig.Addons {
		var prefix string = "├─ "
		if i == len(projectConfig.Addons)-1 {
			prefix = "└─ "
		}
		var line []string = []string{prefix + string(folder) + "/", addon.Id, "(" + addon.Asset + ")"}
		fmt.Fprintln(writer, strings.Join(line, "\t"))
		i++
	}
	fmt.Fprintln(writer)
	writer.Flush()
}

func SearchAddon(search string) {
	var result AssetResult
	if err := doGet("/asset?filter="+search+"&godot_version=3.4", &result); err != nil {
		return
	}
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(writer, "[id]\t[user]\t[title]")
	for _, addon := range result.Result {
		fmt.Fprintln(writer, strings.Join([]string{addon.AssetId, addon.Author, addon.Title}, "\t"))
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
