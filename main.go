package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	arguments := os.Args
	flavor := arguments[1]
	image := arguments[2]

	currentDirectory, _ := os.Getwd()

	var command string

	switch flavor {
	case "latte":
		command = fmt.Sprintf("cd %s ; magick %s flavors/latte-hald-clut.png -hald-clut %s-%s", currentDirectory, image, flavor, image)
	case "frappe":
		command = fmt.Sprintf("cd %s ; magick %s flavors/frappe-hald-clut.png -hald-clut %s-%s", currentDirectory, image, flavor, image)
	case "macchiato":
		command = fmt.Sprintf("cd %s ; magick %s flavors/macchiato-hald-clut.png -hald-clut %s-%s", currentDirectory, image, flavor, image)
	case "mocha":
		command = fmt.Sprintf("cd %s ; magick %s flavors/mocha-hald-clut.png -hald-clut %s-%s", currentDirectory, image, flavor, image)
	default:
		command = ""
	}

	if command == "" {
		fmt.Println("Flavor not found")
		return
	}

	cmd := exec.Command("bash", "-c", command)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Successfully generated image with %s flavor\n", flavor)
}
