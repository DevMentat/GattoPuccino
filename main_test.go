package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	testCases := []struct {
		description string
		args        []string
	}{
		{"Valid input", []string{"gattopuccino", "latte", "test-image.jpg"}},
		{"Invalid flavor", []string{"gattopuccino", "invalid-flavor", "test-image.jpg"}},
		{"Missing input arguments", []string{"gattopuccino"}},
		{"Missing image argument", []string{"gattopuccino", "latte"}},
		{"Missing flavor argument", []string{"gattopuccino", "test-image.jpg"}},
		{"Missing image file", []string{"gattopuccino", "latte", "missing-image.jpg"}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			os.Args = tc.args
			main()
		})
	}
}
