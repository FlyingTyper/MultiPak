package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/FlyingTyper/multipak/pkg/cli"
)

func main() {
	// Parse command-line flags
	searchCommand := flag.NewFlagSet("search", flag.ExitOnError)
	packageName := searchCommand.String("package", "", "Package name to search")

	// Parse command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: MultiPak [options]")
		os.Exit(1)
	}

	// Handle commands
	switch os.Args[1] {
	case "/S,S":
		// Parse search command flags for Snap
		searchCommand.Parse(os.Args[2:])
		if *packageName == "" {
			fmt.Println("Error: Please provide a package name to search.")
			os.Exit(1)
		}
		cli.SearchSnapPackage(*packageName)

	case "/S,F":
		// Parse search command flags for Flatpak
		searchCommand.Parse(os.Args[2:])
		if *packageName == "" {
			fmt.Println("Error: Please provide a package name to search.")
			os.Exit(1)
		}
		cli.SearchFlatpakPackage(*packageName)

	case "/S,A":
		// Parse search command flags for Arch AUR
		searchCommand.Parse(os.Args[2:])
		if *packageName == "" {
			fmt.Println("Error: Please provide a package name to search.")
			os.Exit(1)
		}
		cli.SearchAURPackage(*packageName)

	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}
