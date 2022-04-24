package packagemanager

import (
	"encoding/json"
	swagger "gibson/package_manager/client"
	"os"
	"path/filepath"
)

const (
	PACKAGE_MANAGER_V      string = "1.1.0"
	ASSETS_FOLDER          string = "addons"
	ARCHIVE_NAME_DELIMITER string = "$"
)

var (
	CACHE_FOLDER string       = filepath.Join("gibson", "addons")
	cachedConfig CachedConfig = CachedConfig{
		Version:      PACKAGE_MANAGER_V,
		CachedAssets: map[string]swagger.AssetDetails{},
	}
	projectConfig ProjectConfig = ProjectConfig{
		Version: PACKAGE_MANAGER_V,
		Assets:  map[string]GibsonAsset{},
	}
)

func initConfig() {
	folder, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	CACHE_FOLDER = filepath.Join(folder, CACHE_FOLDER)

	if _, err := os.Stat(CACHE_FOLDER); err != nil {
		os.MkdirAll(CACHE_FOLDER, os.ModePerm)
	}

	if err = loadCachedConfig(); err != nil {
		if err = saveCachedConfig(); err != nil {
			panic(err)
		}
	}

	if err = loadConfig(); err != nil {
		if err = saveConfig(); err != nil {
			panic(err)
		}
	}
}

func loadCachedConfig() error {
	var config CachedConfig
	content, err := os.Open(filepath.Join(CACHE_FOLDER, "gibson.json"))
	if err != nil {
		return err
	}
	defer content.Close()

	err = json.NewDecoder(content).Decode(&config)
	cachedConfig = config
	return nil
}

func saveCachedConfig() error {
	marshal, err := json.Marshal(cachedConfig)
	err = os.WriteFile(filepath.Join(CACHE_FOLDER, "gibson.json"), marshal, os.ModePerm)
	return err
}

func addCachedAsset(assetDetails swagger.AssetDetails) {
	cachedConfig.CachedAssets[assetDetails.Author+"/"+assetDetails.Title] = assetDetails
	saveCachedConfig()
}

func removeCachedAsset(fullName string) {
	delete(cachedConfig.CachedAssets, fullName)
	saveCachedConfig()
}

func loadConfig() error {
	var config ProjectConfig
	content, err := os.Open("gibson.json")
	if err != nil {
		return err
	}
	defer content.Close()

	err = json.NewDecoder(content).Decode(&config)
	projectConfig = config
	return nil
}

func saveConfig() error {
	marshal, err := json.Marshal(projectConfig)
	err = os.WriteFile("gibson.json", marshal, os.ModePerm)
	return err
}

func addAsset(currentFolder string, assetDetails swagger.AssetDetails) {
	projectConfig.Assets[currentFolder] = GibsonAsset{
		Id:       assetDetails.AssetId,
		FullName: assetDetails.Author + "/" + assetDetails.Title,
		Version:  assetDetails.VersionString,
	}
	saveConfig()
}

func removeAsset(assetId string) {
	for currentFolder, gibsonAsset := range projectConfig.Assets {
		if gibsonAsset.Id == assetId || currentFolder == assetId {
			delete(projectConfig.Assets, currentFolder)
			break
		}
	}
	saveConfig()
}
