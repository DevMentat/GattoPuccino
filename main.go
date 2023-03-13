package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Get the command line arguments
	arguments := os.Args
	flavor := arguments[1] // The flavor of the coffee
	image := arguments[2]  // The name of the input image

	// Get the current working directory
	currentDirectory, _ := os.Getwd()

	// Set the output directory name
	outputDir := "output"

	var command string

	// Build the command based on the chosen flavor
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
		command = ""
	}

	// Check if the flavor is recognized
	if command == "" {
		fmt.Println("Flavor not found")
		return
	}

	// Create the output directory if it does not exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Execute the command to generate the output image
	cmd := exec.Command("bash", "-c", command)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Print a message indicating success and the location of the output image
	fmt.Printf("Successfully generated image with %s flavor in %s/%s-%s\n", flavor, outputDir, flavor, image)
}
