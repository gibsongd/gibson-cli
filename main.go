package main

import (
	"flag"
	"fmt"
	packagemanager "gibson/package_manager"
	"os"
)

var Version string = "1.1.0"

func main() {
	pmCmd := flag.NewFlagSet("pm", flag.ExitOnError)

	var searchAddon string
	var installAddon string
	var uninstallAddon string
	var clearCached bool
	var listAddons bool

	pmCmd.StringVar(&searchAddon, "search", "", "Search for an addon")
	pmCmd.StringVar(&installAddon, "install", "", "Install a new addon")
	pmCmd.StringVar(&uninstallAddon, "uninstall", "", "Uninstall a new addon")
	pmCmd.BoolVar(&listAddons, "list", false, "List all addons installed in the current project")
	pmCmd.BoolVar(&clearCached, "clear", false, "Clear cached addon")

	versionPtr := flag.Bool("version", false, "Print gibson-cli version")
	flag.Parse()

	if *versionPtr {
		fmt.Println(Version)
		os.Exit(0)
	}
	if len(os.Args) < 2 {
		flag.Usage()
		pmCmd.Usage()
		os.Exit(0)
	}
	switch os.Args[1] {
	case "pm":
		packagemanager.Init()

		pmCmd.Parse(os.Args[2:])

		if listAddons {
			packagemanager.ListAddons()
			os.Exit(0)
		}
		if searchAddon != "" {
			packagemanager.SearchAddon(searchAddon)
			os.Exit(0)
		}
		if installAddon == "" && uninstallAddon == "" {
			pmCmd.Usage()
			os.Exit(0)
		}
		if installAddon != "" {
			if installAddon == "." {
				packagemanager.InstallByConfig(clearCached)
			} else {
				packagemanager.InstallAddon(installAddon, clearCached)
			}
		}
		if uninstallAddon != "" {
			packagemanager.UninstallAddon(uninstallAddon, clearCached)
		}
	default:
		os.Exit(0)
	}

}
