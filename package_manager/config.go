package packagemanager

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var cachedConfig CachedConfig = CachedConfig{
	CachedAddons: map[AssetName]Addon{},
}
var projectConfig ProjectConfig = ProjectConfig{
	Addons: map[Folder]GibsonAddon{},
}

var cacheFolder string = filepath.Join("gibson", "addons")
var addonsFolder string = "addons"

func initConfig() {
	folder, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	cacheFolder = filepath.Join(folder, cacheFolder)

	if _, err := os.Stat(cacheFolder); err != nil {
		os.MkdirAll(cacheFolder, os.ModePerm)
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
	content, err := os.Open(filepath.Join(cacheFolder, "gibson.json"))
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
	err = os.WriteFile(filepath.Join(cacheFolder, "gibson.json"), marshal, os.ModePerm)
	return err
}

func addCachedAsset(asset string, addon Addon) {
	if len(cachedConfig.CachedAddons) == 0 {
		cachedConfig.CachedAddons = map[AssetName]Addon{AssetName(asset): addon}
	} else {
		cachedConfig.CachedAddons[AssetName(asset)] = addon
	}
	saveCachedConfig()
}

func removeCachedAsset(asset string) {
	delete(cachedConfig.CachedAddons, AssetName(asset))
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

func addAddon(folder Folder, addon GibsonAddon) {
	if len(projectConfig.Addons) == 0 {
		projectConfig.Addons = map[Folder]GibsonAddon{folder: addon}
	} else {
		projectConfig.Addons[folder] = addon
	}
	saveConfig()
}

func removeAddon(folder Folder) {
	delete(projectConfig.Addons, folder)
	saveConfig()
}
