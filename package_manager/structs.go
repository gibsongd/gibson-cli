package packagemanager

type Addon struct {
	AssetId        string `json:"asset_id"`
	Author         string `json:"author"`
	Title          string `json:"title"`
	Version        string `json:"version"`
	DownloadCommit string `json:"download_commit"`
	DownloadUrl    string `json:"download_url"`
}

type GibsonAddon struct {
	Asset string `json:"asset"`
	Id    string `json:"id"`
}

type AssetResult struct {
	Result []Addon `json:"result"`
}

type Folder string
type ProjectConfig struct {
	Version string                 `json:"version"`
	Addons  map[Folder]GibsonAddon `json:"addons"`
}

type AssetName string
type CachedConfig struct {
	Version      string              `json:"version"`
	CachedAddons map[AssetName]Addon `json:"cachedAddons"`
}
