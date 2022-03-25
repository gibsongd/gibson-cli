package packagemanager

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/theckman/yacspin"
)

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

var baseUrl string = "https://godotengine.org/asset-library/api"

var downloadSpinner *yacspin.Spinner
var getSpinner *yacspin.Spinner

var colorYellow string = "\033[33m"
var colorReset string = "\033[0m"

func init() {

	getSpinner, _ = yacspin.New(
		yacspin.Config{
			Frequency:         100 * time.Millisecond,
			CharSet:           yacspin.CharSets[59],
			Prefix:            "[gibson-cli] ",
			StopCharacter:     "✓",
			StopColors:        []string{"fgGreen"},
			StopFailCharacter: "✗",
			StopFailColors:    []string{"fgRed"},
		},
	)

	downloadSpinner, _ = yacspin.New(
		yacspin.Config{
			Frequency:         100 * time.Millisecond,
			CharSet:           yacspin.CharSets[59],
			Prefix:            "[gibson-cli] ",
			StopCharacter:     "✓",
			StopColors:        []string{"fgGreen"},
			StopFailCharacter: "✗",
			StopFailColors:    []string{"fgRed"},
		},
	)

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

func unzip(src string, dest string, skipName string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		fpath = strings.Replace(fpath, "-"+skipName, "", 1)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func doGet(url string, target interface{}) (int, error) {
	resp, err := http.Get(baseUrl + url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, json.NewDecoder(resp.Body).Decode(target)
}

func spinnerMessage(message string, spinner *yacspin.Spinner) {
	spinner.Message(" " + message)
}

func startSpinner(message string, spinner *yacspin.Spinner) {
	spinner.Start()
	spinnerMessage(message, spinner)
}

func stopSpinner(message string, spinner *yacspin.Spinner) {
	spinner.StopMessage(" " + message)
	spinner.Stop()
}

func failSpinner(message string, spinner *yacspin.Spinner) {
	spinner.StopFailMessage(" " + message)
	spinner.StopFail()
}

func formatAsset(asset string) string {
	return "`" + colorYellow + asset + colorReset + "`"
}

func InstallByAuthor(asset string) {

	tempPack := strings.Split(asset, "/")
	var author string = tempPack[0]
	var assetName string = tempPack[1]

	var toDownloadAddon Addon = Addon{}
	var assetResult AssetResult = AssetResult{}

	startSpinner("looking for "+formatAsset(asset), getSpinner)
	_, err := doGet("/asset?user="+author+"&godot_version=3.4", &assetResult)
	if err != nil {
		failSpinner("fetching failed, reason: "+err.Error(), getSpinner)
	} else {
		if len(assetResult.Result) < 1 {
			failSpinner("couldn't find assets related to @"+author, getSpinner)
		}
		for _, addon := range assetResult.Result {
			if addon.Title == assetName {
				toDownloadAddon = addon
				spinnerMessage(formatAsset(asset)+" found!", getSpinner)
				break
			}
		}
		if toDownloadAddon.Asset_id == "" {
			failSpinner("couldn't find "+formatAsset(asset), getSpinner)
		} else {
			stopSpinner("starting download...", getSpinner)
			startSpinner("downloading "+toDownloadAddon.Title, downloadSpinner)

			var code int
			code, err = doGet("/asset/"+toDownloadAddon.Asset_id, &toDownloadAddon)
			if code > 400 {
				failSpinner("download failed, reason: "+toDownloadAddon.Error, downloadSpinner)
				return
			}

			err = downloadFile(toDownloadAddon.Download_commit+".zip", toDownloadAddon.Download_url)
			if err != nil {
				failSpinner("download failed, reason: "+err.Error(), downloadSpinner)
				return
			} else {
				stopSpinner(formatAsset(asset)+" downloaded!", downloadSpinner)
			}

			unzip(toDownloadAddon.Download_commit+".zip", "addons/", toDownloadAddon.Download_commit)
		}
	}

}

func InstallById(assetId string) {
	var toDownloadAddon Addon = Addon{}
	startSpinner("downloading "+toDownloadAddon.Title, downloadSpinner)

	code, err := doGet("/asset/"+assetId, &toDownloadAddon)
	if code > 400 {
		failSpinner("download failed, reason: "+toDownloadAddon.Error, downloadSpinner)
		return
	}

	err = downloadFile(toDownloadAddon.Download_commit+".zip", toDownloadAddon.Download_url)
	if err != nil {
		failSpinner("download failed, reason: "+err.Error(), downloadSpinner)
	} else {
		stopSpinner(formatAsset(toDownloadAddon.Author+"/"+toDownloadAddon.Title)+" downloaded!", downloadSpinner)
	}
}
