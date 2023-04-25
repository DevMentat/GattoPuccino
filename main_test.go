package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Save original arguments and restore them after tests
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	testCases := []struct {
		args []string
	}{
		{[]string{"gattopuccino", "latte", "test-image.jpg"}},
		{[]string{"gattopuccino", "invalid-flavor", "test-image.jpg"}},
		{[]string{"gattopuccino"}},
		{[]string{"gattopuccino", "latte"}},
		{[]string{"gattopuccino", "test-image.jpg"}},
		{[]string{"gattopuccino", "latte", "missing-image.jpg"}},
	}

	for _, tc := range testCases {
		os.Args = tc.args
		main()
	}
}
