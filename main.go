package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: gattopuccino <flavor> <image>")
		return
	}
	flavor := os.Args[1]
	image := os.Args[2]

	if _, err := os.Stat(image); os.IsNotExist(err) {
		fmt.Printf("Error: image file '%s' not found\n", image)
		return
	}

	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	outputDir := "output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	flavors := map[string]string{
		"latte":     "flavors/latte-hald-clut.png",
		"frappe":    "flavors/frappe-hald-clut.png",
		"macchiato": "flavors/macchiato-hald-clut.png",
		"mocha":     "flavors/mocha-hald-clut.png",
	}

	lookupFlavor, ok := flavors[flavor]
	if !ok {
		fmt.Println("Error: flavor not recognized")
		fmt.Println("Usage: gattopuccino <flavor> <image>")
		return
	}

	outputPath := filepath.Join(outputDir, fmt.Sprintf("%s-%s", flavor, image))
	cmd := exec.Command("magick", image, lookupFlavor, "-hald-clut", outputPath)
	cmd.Dir = currentDirectory

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Successfully generated image with %s flavor in %s\n", flavor, outputPath)
}
