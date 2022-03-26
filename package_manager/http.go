package packagemanager

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

const BASE_URL string = "https://godotengine.org/asset-library/api"

func doGet(url string, target interface{}) (int, error) {
	resp, err := http.Get(BASE_URL + url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, json.NewDecoder(resp.Body).Decode(target)
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
