package packagemanager

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func formatAsset(asset string) string {
	return "`" + colorYellow + asset + colorReset + "`"
}

func unzip(src string, dest string) (string, error) {
	var packageName string

	r, err := zip.OpenReader(src)
	if err != nil {
		return packageName, err
	}
	defer r.Close()

	var basePath string

	for i, f := range r.File {

		if i == 0 {
			basePath = f.Name + addonsFolder
			continue
		}

		if strings.Contains(f.Name, "/addons/") {
			src, err := f.Open()
			if err != nil {
				return packageName, err
			}

			var outFileName string = strings.Replace(f.Name, basePath, "", 1)
			if packageName == "" {
				packageName = strings.Split(outFileName, "/")[1]
			}
			var outFilePath string = filepath.Join(dest, outFileName)

			if f.FileInfo().IsDir() {
				os.Mkdir(outFilePath, os.ModePerm)
				continue
			}

			if err := os.MkdirAll(filepath.Dir(outFilePath), os.ModePerm); err != nil {
				return packageName, err
			}

			out, err := os.OpenFile(outFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return packageName, err
			}

			_, err = io.Copy(out, src)
			if err != nil {
				return packageName, err
			}

			src.Close()
			out.Close()
		}
	}
	return packageName, nil
}
