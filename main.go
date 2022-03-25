package main

import (
	"flag"
	"fmt"
	packagemanager "gibson/gisbon-cli/package-manager"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Gibson CLI ===")

	pmCmd := flag.NewFlagSet("pm", flag.ExitOnError)

	var asset string
	pmCmd.StringVar(&asset, "install", "", "Install a new asset")

	fmt.Println(os.Args, len(os.Args))

	if len(os.Args) < 2 {
		fmt.Println("expected 'pm' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "pm":
		pmCmd.Parse(os.Args[2:])
		if strings.Contains(asset, "/") {
			packagemanager.InstallByAuthor(asset)
		} else {
			packagemanager.InstallById(asset)
		}
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

}
