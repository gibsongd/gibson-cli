package packagemanager

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

const BASE_URL string = "https://godotengine.org/asset-library/api"

type HttpError struct {
	Error string
}

func doGet(url string, target interface{}) error {
	resp, err := http.Get(BASE_URL + url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var httpError HttpError
		json.NewDecoder(resp.Body).Decode(httpError)
		err = errors.New(httpError.Error)
	} else {
		err = json.NewDecoder(resp.Body).Decode(target)
	}

	return err
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
