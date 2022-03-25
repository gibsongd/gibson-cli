package packagemanager

import (
	"archive/zip"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func formatAsset(asset string) string {
	return "`" + colorYellow + asset + colorReset + "`"
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return err
}

func unzip(src string, dest string) error {

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	var basePath string

	for i, f := range r.File {

		if i == 0 {
			basePath = f.Name + unzipTarget
			continue
		}

		if strings.Contains(f.Name, "/addons/") {
			src, err := f.Open()
			if err != nil {
				return err
			}

			var outFileName string = strings.Replace(f.Name, basePath, "", 1)
			var outFilePath string = filepath.Join(dest, outFileName)

			if f.FileInfo().IsDir() {
				os.Mkdir(outFilePath, os.ModePerm)
				continue
			}

			if err := os.MkdirAll(filepath.Dir(outFilePath), os.ModePerm); err != nil {
				return err
			}

			out, err := os.OpenFile(outFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			_, err = io.Copy(out, src)
			if err != nil {
				return err
			}

			src.Close()
			out.Close()
		}
	}
	return nil
}

func doGet(url string, target interface{}) (int, error) {
	resp, err := http.Get(baseUrl + url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, json.NewDecoder(resp.Body).Decode(target)
}
