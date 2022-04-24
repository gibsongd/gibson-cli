package packagemanager

import swagger "gibson/package_manager/client"

type GibsonAsset struct {
	Id       string `json:"id"`
	FullName string `json:"fullName"`
	Version  string `json:"version"`
}

type ProjectConfig struct {
	Version string                 `json:"version"`
	Assets  map[string]GibsonAsset `json:"assets"` // map[localFolder]{id, fullName, version}
}

type CachedConfig struct {
	Version      string                          `json:"version"`
	CachedAssets map[string]swagger.AssetDetails `json:"cachedAssets"` // map[id]asset
}
