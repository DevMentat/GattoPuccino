package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Test case for a successful execution with valid input
	os.Args = []string{"gattopuccino", "latte", "test-image.jpg"}
	main()

	// Test case for an invalid flavor
	os.Args = []string{"gattopuccino", "invalid-flavor", "test-image.jpg"}
	main()

	// Test case for missing input arguments
	os.Args = []string{"gattopuccino"}
	main()

	// Test case for missing image argument
	os.Args = []string{"gattopuccino", "latte"}
	main()

	// Test case for missing flavor argument
	os.Args = []string{"gattopuccino", "test-image.jpg"}
	main()

	// Test case for a failed execution due to a missing image file
	os.Args = []string{"gattopuccino", "latte", "missing-image.jpg"}
	main()
}
