package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Gattopuccino by Devmentat
	// Get the command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: gattopuccino <flavor> <image>")
		return
	}
	flavor := os.Args[1] // The flavor of the coffee
	image := os.Args[2]  // The name of the input image

	// Get the current working directory
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Set the output directory name
	outputDir := "output"

	// Create the output directory if it does not exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Check if the flavor is recognized
	var command string
	switch flavor {
	case "latte":
		// Use the "latte" flavor
		command = fmt.Sprintf("cd %s ; magick %s flavors/latte-hald-clut.png -hald-clut %s/%s-%s", currentDirectory, image, outputDir, flavor, image)
	case "frappe":
		// Use the "frappe" flavor
		command = fmt.Sprintf("cd %s ; magick %s flavors/frappe-hald-clut.png -hald-clut %s/%s-%s", currentDirectory, image, outputDir, flavor, image)
	case "macchiato":
		// Use the "macchiato" flavor
		command = fmt.Sprintf("cd %s ; magick %s flavors/macchiato-hald-clut.png -hald-clut %s/%s-%s", currentDirectory, image, outputDir, flavor, image)
	case "mocha":
		// Use the "mocha" flavor
		command = fmt.Sprintf("cd %s ; magick %s flavors/mocha-hald-clut.png -hald-clut %s/%s-%s", currentDirectory, image, outputDir, flavor, image)
	default:
		// Flavor not recognized
		fmt.Println("Error: flavor not recognized")
		fmt.Println("Usage: gattopuccino <flavor> <image>")
		return
	}

	// Build the command based on the chosen flavor
	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = currentDirectory

	// Execute the command to generate the output image
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Print a message indicating success and the location of the output image
	fmt.Printf("Successfully generated image with %s flavor in %s/%s-%s\n", flavor, outputDir, flavor, image)
}
