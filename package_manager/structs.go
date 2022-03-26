package packagemanager

type Addon struct {
	Asset_id        string
	Author          string
	Title           string
	Version         string
	Download_commit string
	Download_url    string

	Error string
}

type AssetResult struct {
	Result []Addon
}
