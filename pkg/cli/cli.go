package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Function to search for a package in Snap
func SearchSnapPackage(packageName string) {
	cmd := exec.Command("snap", "find", packageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error searching Snap package:", err)
		return
	}

	// Extract and print relevant information from Snap search results
	printResultsFromSnap(string(output))
}

// Function to search for a package in Flatpak
func SearchFlatpakPackage(packageName string) {
	cmd := exec.Command("flatpak", "search", packageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error searching Flatpak package:", err)
		return
	}

	// Extract and print relevant information from Flatpak search results
	printResultsFromFlatpak(string(output))
}

// Function to search for a package in Arch AUR
func SearchAURPackage(packageName string) {
	cmd := exec.Command("yay", "-Ss", packageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error searching AUR package:", err)
		return
	}

	// Extract and print relevant information from AUR search results
	printResultsFromAUR(string(output))
}

// Function to print package names and descriptions from Snap search results
func printResultsFromSnap(searchResults string) {
	lines := strings.Split(searchResults, "\n")
	for _, line := range lines {
		// Extract relevant information from Snap search results (modify as needed)
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) >= 3 {
				name := strings.TrimSpace(parts[0])
				description := strings.TrimSpace(parts[2])
				fmt.Printf("Name: %s\nDescription: %s\n", name, description)
			}
		}
	}
}

// Function to print package names and descriptions from Flatpak search results
func printResultsFromFlatpak(searchResults string) {
	lines := strings.Split(searchResults, "\n")
	for _, line := range lines {
		// Extract relevant information from Flatpak search results (modify as needed)
		if strings.Contains(line, "Application ID") {
			parts := strings.Split(line, " ")
			if len(parts) >= 3 {
				name := strings.TrimSpace(parts[2])
				fmt.Printf("Name: %s\n", name)
			}
		} else if strings.Contains(line, "Description") {
			description := strings.TrimSpace(strings.TrimPrefix(line, "Description:"))
			fmt.Printf("Description: %s\n", description)
		}
	}
}

// Function to print package names and descriptions from AUR search results
func printResultsFromAUR(searchResults string) {
	lines := strings.Split(searchResults, "\n")
	for _, line := range lines {
		// Extract relevant information from AUR search results (modify as needed)
		if strings.HasPrefix(line, "aur/") {
			parts := strings.Fields(line)
			if len(parts) >= 5 {
				name := strings.TrimSpace(parts[1])
				description := strings.TrimSpace(strings.Join(parts[4:], " "))
				fmt.Printf("Name: %s\nDescription: %s\n", name, description)
			}
		}
	}
}
