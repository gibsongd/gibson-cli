package main

import (
	"flag"
	"fmt"
	packagemanager "gibson/gisbon-cli/package-manager"
	"os"
)

func main() {
	fmt.Println("=== Gibson CLI ===")

	pmCmd := flag.NewFlagSet("pm", flag.ExitOnError)

	var installAsset string
	var uninstallAsset string
	var clearCached bool

	pmCmd.StringVar(&installAsset, "install", "", "Install a new asset")
	pmCmd.StringVar(&uninstallAsset, "uninstall", "", "Uninstall a new asset")
	pmCmd.BoolVar(&clearCached, "clear", false, "Clear cached asset")

	if len(os.Args) < 2 {
		fmt.Println("expected 'pm' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "pm":
		pmCmd.Parse(os.Args[2:])
		if installAsset != "" {
			packagemanager.InstallAddon(installAsset, clearCached)
		}
		if uninstallAsset != "" {
			packagemanager.UninstallAsset(uninstallAsset)
		}
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

}
