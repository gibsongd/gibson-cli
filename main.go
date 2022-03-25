package main

import (
	"flag"
	"fmt"
	packagemanager "gibson/package-manager"
	"os"
)

var version string = "1.0.0"

func main() {
	pmCmd := flag.NewFlagSet("pm", flag.ExitOnError)

	var installAsset string
	var uninstallAsset string
	var clearCached bool

	pmCmd.StringVar(&installAsset, "install", "", "Install a new asset")
	pmCmd.StringVar(&uninstallAsset, "uninstall", "", "Uninstall a new asset")
	pmCmd.BoolVar(&clearCached, "clear", false, "Clear cached asset")

	versionPtr := flag.Bool("version", false, "Print gibson-cli version")
	flag.Parse()

	if *versionPtr {
		fmt.Println(getVersion())
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println(getUsage())
		os.Exit(1)
	}

	switch os.Args[1] {
	case "pm":
		pmCmd.Parse(os.Args[2:])
		if installAsset == "" && uninstallAsset == "" {
			fmt.Println(packagemanager.GetUsage())
			os.Exit(1)
		}
		if installAsset != "" {
			packagemanager.InstallAddon(installAsset, clearCached)
		}
		if uninstallAsset != "" {
			packagemanager.UninstallAsset(uninstallAsset, clearCached)
		}
	default:
		fmt.Println(getUsage())
		os.Exit(1)
	}

}

func getVersion() string {
	return version
}

func getUsage() string {
	return "Usage:\n\n" +
		"gibson pm    Package Manager\n" +
		"\n" +
		"All commands:\n\n" +
		"    pm, version, help\n"
}
