package config

import (
	"os"
	"testing"
)

// go test scans every file ending in _test.go/
// and the function which we have test must start with Test

func TestLoad(t *testing.T) {

	// Temporary YAML content for testing.
	data := []byte("brute_force:\n  threshold: 5\n  window: 30s\n")

	// Create a temporary config file.
	file := "test_config.yml"

	err := os.WriteFile(file, data, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Delete the temporary file after the test. runs at the end of function
	defer os.Remove(file)

	// Load the configuration.
	cfg, err := Load(file)

	if err != nil {
		t.Fatal(err)
	}

	// Check threshold.
	if cfg.BruteForce.Threshold != 5 {
		t.Fatalf(
			"expected threshold 5, got %d",
			cfg.BruteForce.Threshold,
		)
	}

	// Check window.
	if cfg.BruteForce.Window != "30s" {
		t.Fatalf(
			"expected window 30s, got %s",
			cfg.BruteForce.Window,
		)
	}
}
