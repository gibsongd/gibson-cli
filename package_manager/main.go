package packagemanager

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var cacheFolder string = "gibson/addons"
var unzipTarget string = "addons"

func initPM() {
	folder, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	cacheFolder = filepath.Join(folder, cacheFolder)
}

func init() {
	initPM()
	initSpinners()
}

func GetUsage() string {
	return "Usage of pm:\n" +
		"-clear\n" +
		"	Clear cached asset\n" +
		"-install string\n" +
		"	Install a new asset\n" +
		"-uninstall string\n" +
		"	Uninstall a new asset\n"
}

func clearCache(asset string) {
	var assetPath string = filepath.Join(cacheFolder, asset)
	startSpinner("clearing cached "+formatAsset(asset), cacheSpinner)
	if err := os.RemoveAll(assetPath); err == nil {
		stopSpinner(formatAsset(asset)+" cache cleared", cacheSpinner)
	} else {
		failSpinner("could not clear cached "+formatAsset(asset)+", reason: "+err.Error(), cacheSpinner)
	}
}

func lookForCachedAsset(asset string) (string, error) {
	if strings.Contains(asset, "/") {
		var assetPath string = filepath.Join(cacheFolder, asset)
		if _, err := os.Stat(assetPath); err == nil {
			fs, err := ioutil.ReadDir(assetPath)
			if err != nil {
				return "", err
			}
			sort.SliceStable(fs, func(i, j int) bool {
				return fs[i].Name() > fs[j].Name()
			})
			return filepath.Join(assetPath, fs[0].Name()), nil
		}
	} else {
		// maybe the id of each asset should be stored in a file under "gibson/addons"
	}
	return "", nil
}

func installAsset(asset string, archive string) {
	startSpinner("installing "+formatAsset(asset), installSpinner)
	err := unzip(archive, unzipTarget)
	if err != nil {
		failSpinner("could not install "+formatAsset(asset)+", reason: "+err.Error(), installSpinner)
		return
	}
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

	if clearCached {
		clearCache(asset)
	} else {
		if cached, err := lookForCachedAsset(asset); err == nil && cached != "" {
			stopSpinner("found cached version of "+formatAsset(asset), findSpinner)
			installAsset(asset, cached)
			return
		}
	}

	_, err := doGet("/asset?user="+author+"&godot_version=3.4", &assetResult)
	if err != nil {
		failSpinner("fetching failed, reason: "+err.Error(), findSpinner)
		return
	}

	// If there's a list, look for the asset
	if len(assetResult.Result) < 1 {
		failSpinner("couldn't find assets related to @"+author, findSpinner)
	}
	for _, addon := range assetResult.Result {
		if addon.Title == assetName {
			toDownloadAddon = addon
			spinnerMessage(formatAsset(asset)+" found!", findSpinner)
			break
		}
	}

	// If there's an asset, download it and install it
	if toDownloadAddon.Asset_id == "" {
		failSpinner("couldn't find "+formatAsset(asset), findSpinner)
		return
	}
	stopSpinner(formatAsset(asset)+" fetched!", findSpinner)

	installById(toDownloadAddon.Asset_id)

}

func installById(assetId string) {
	var toDownloadAddon Addon = Addon{Asset_id: assetId}

	startSpinner("retrieving "+formatAsset(assetId)+" info", getSpinner)

	code, err := doGet("/asset/"+assetId, &toDownloadAddon)
	if code > 400 {
		failSpinner("could not find "+formatAsset(assetId)+", reason: "+toDownloadAddon.Error, getSpinner)
		return
	}
	stopSpinner(formatAsset(assetId)+" info retrieved!", getSpinner)

	var asset string = toDownloadAddon.Author + "/" + toDownloadAddon.Title
	var assetFolder string = filepath.Join(cacheFolder, asset)
	os.MkdirAll(assetFolder, os.ModePerm)

	startSpinner("downloading "+formatAsset(asset), downloadSpinner)

	var archive string = filepath.Join(assetFolder, toDownloadAddon.Version+"_"+toDownloadAddon.Download_commit+".zip")
	err = downloadFile(archive, toDownloadAddon.Download_url)
	if err != nil {
		failSpinner("download failed, reason: "+err.Error(), downloadSpinner)
		return
	}
	stopSpinner(formatAsset(asset)+" downloaded!", downloadSpinner)

	installAsset(asset, archive)
}

func InstallAddon(asset string, clearCached bool) {
	if strings.Contains(asset, "/") {
		installByAuthor(asset, clearCached)
	} else {
		installById(asset)
	}
}

func UninstallAsset(asset string, clearCached bool) {
	startSpinner("uninstalling "+formatAsset(asset), findSpinner)
	if clearCached {
		clearCache(asset)
	}
	// remove local addon
	stopSpinner(formatAsset(asset)+" uninstalled", findSpinner)

}
